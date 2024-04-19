import { isUndefined, uniq } from "lodash-es";
import { defineStore } from "pinia";
import { computed, ref, unref, watch, watchEffect } from "vue";
import { projectServiceClient } from "@/grpcweb";
import type { ComposedDatabase, MaybeRef } from "@/types";
import { ALL_USERS_USER_EMAIL, PresetRoleType } from "@/types";
import { IamPolicy } from "@/types/proto/v1/iam_policy";
import {
  hasWorkspacePermissionV2,
  isDeveloperOfProjectV1,
  isOwnerOfProjectV1,
  isViewerOfProjectV1,
} from "@/utils";
import { convertFromExpr } from "@/utils/issue/cel";
import { useCurrentUserV1 } from "../auth";
import { useProjectV1Store } from "./project";

export const useProjectIamPolicyStore = defineStore(
  "project-iam-policy",
  () => {
    const policyMap = ref(new Map<string, IamPolicy>());
    const requestCache = new Map<string, Promise<IamPolicy>>();

    const fetchProjectIamPolicy = async (
      project: string,
      skipCache = false
    ) => {
      if (!skipCache) {
        const cache = requestCache.get(project);
        if (cache) {
          return cache;
        }
      }

      const request = projectServiceClient
        .getIamPolicy({
          project,
        })
        .then((policy) => {
          policyMap.value.set(project, policy);
          return policy;
        });
      requestCache.set(project, request);
      return request;
    };

    const batchFetchIamPolicy = async (projectList: string[]) => {
      const response = await projectServiceClient.batchGetIamPolicy({
        scope: "projects/-",
        names: projectList,
      });
      response.policyResults.forEach(({ policy, project }) => {
        if (policy) {
          policyMap.value.set(project, policy);
        }
      });
    };

    const updateProjectIamPolicy = async (
      project: string,
      policy: IamPolicy
    ) => {
      policy.bindings.forEach((binding) => {
        if (binding.members) {
          binding.members = uniq(binding.members);
        }
      });
      const updated = await projectServiceClient.setIamPolicy({
        project,
        policy,
      });
      policyMap.value.set(project, updated);
    };

    const getProjectIamPolicy = (project: string) => {
      return policyMap.value.get(project) ?? IamPolicy.fromJSON({});
    };

    const getOrFetchProjectIamPolicy = async (project: string) => {
      if (!policyMap.value.has(project)) {
        await fetchProjectIamPolicy(project);
      }
      return getProjectIamPolicy(project);
    };

    const batchGetOrFetchProjectIamPolicy = async (
      projectList: string[],
      skipCache = false
    ) => {
      if (skipCache) {
        await batchFetchIamPolicy(projectList);
      } else {
        // BatchFetch policies that missing in the local map.
        const missingProjectList = projectList.filter(
          (project) => !policyMap.value.has(project)
        );
        if (missingProjectList.length > 0) {
          await batchFetchIamPolicy(missingProjectList);
        }
      }
      return projectList.map(getProjectIamPolicy);
    };

    return {
      policyMap,
      getProjectIamPolicy,
      fetchProjectIamPolicy,
      getOrFetchProjectIamPolicy,
      batchGetOrFetchProjectIamPolicy,
      updateProjectIamPolicy,
    };
  }
);

export const useProjectIamPolicy = (project: MaybeRef<string>) => {
  const store = useProjectIamPolicyStore();
  const ready = ref(false);
  watch(
    () => unref(project),
    (project) => {
      ready.value = false;
      store.fetchProjectIamPolicy(project).finally(() => {
        ready.value = true;
      });
    },
    { immediate: true }
  );
  const policy = computed(() => {
    return store.policyMap.get(unref(project)) ?? IamPolicy.fromJSON({});
  });
  return { policy, ready };
};

export const useCurrentUserIamPolicy = () => {
  const iamPolicyStore = useProjectIamPolicyStore();
  const projectStore = useProjectV1Store();
  const currentUser = useCurrentUserV1();

  watchEffect(() => {
    // Fetch all project iam policies.
    Promise.all(
      projectStore.projectList.map((project) =>
        iamPolicyStore.getOrFetchProjectIamPolicy(project.name)
      )
    );
  });

  // hasWorkspaceSuperPrivilege checks whether the current user has the super privilege to access all databases. AKA. Owners and DBAs
  const hasWorkspaceSuperPrivilege = hasWorkspacePermissionV2(
    currentUser.value,
    "bb.projects.list"
  );

  const isProjectOwnerOrDeveloper = (projectName: string): boolean => {
    if (hasWorkspaceSuperPrivilege) {
      return true;
    }

    const policy = iamPolicyStore.policyMap.get(projectName);
    if (!policy) {
      return false;
    }
    return (
      isOwnerOfProjectV1(policy, currentUser.value) ||
      isDeveloperOfProjectV1(policy, currentUser.value)
    );
  };

  const isProjectOwnerOrDeveloperOrViewer = (projectName: string): boolean => {
    if (hasWorkspaceSuperPrivilege) {
      return true;
    }

    const policy = iamPolicyStore.policyMap.get(projectName);
    if (!policy) {
      return false;
    }
    return (
      isProjectOwnerOrDeveloper(projectName) ||
      isViewerOfProjectV1(policy, currentUser.value)
    );
  };

  const allowToChangeDatabaseOfProject = (projectName: string) => {
    return isProjectOwnerOrDeveloper(projectName);
  };

  const allowToQueryDatabaseV1 = (
    database: ComposedDatabase,
    schema?: string,
    table?: string
  ) => {
    if (hasWorkspaceSuperPrivilege) {
      return true;
    }

    // Check if the user has the permission to query the database.
    const policy = database.projectEntity.iamPolicy;
    for (const binding of policy.bindings) {
      if (
        binding.role === PresetRoleType.PROJECT_OWNER &&
        binding.members.find(
          (member) =>
            member === ALL_USERS_USER_EMAIL ||
            member === `user:${currentUser.value.email}`
        )
      ) {
        return true;
      }
      if (
        binding.role === PresetRoleType.PROJECT_QUERIER &&
        binding.members.find(
          (member) =>
            member === ALL_USERS_USER_EMAIL ||
            member === `user:${currentUser.value.email}`
        )
      ) {
        if (binding.parsedExpr?.expr) {
          const conditionExpr = convertFromExpr(binding.parsedExpr.expr);
          if (
            conditionExpr.databaseResources &&
            conditionExpr.databaseResources.length > 0
          ) {
            for (const databaseResource of conditionExpr.databaseResources) {
              if (databaseResource.databaseName === database.name) {
                if (isUndefined(schema) && isUndefined(table)) {
                  return true;
                } else {
                  if (
                    isUndefined(databaseResource.schema) ||
                    (isUndefined(databaseResource.schema) &&
                      isUndefined(databaseResource.table)) ||
                    (databaseResource.schema === schema &&
                      databaseResource.table === table)
                  ) {
                    return true;
                  }
                }
              }
            }
          } else {
            return true;
          }
        } else {
          return true;
        }
      }
    }

    // Otherwise, the user is not allowed to query the database.
    return false;
  };

  const allowToExportDatabaseV1 = (database: ComposedDatabase) => {
    if (hasWorkspaceSuperPrivilege) {
      return true;
    }

    const policy = iamPolicyStore.getProjectIamPolicy(database.project);
    if (!policy) {
      return false;
    }
    const iamPolicyCheckResult = policy.bindings.map((binding) => {
      if (
        binding.role === PresetRoleType.PROJECT_OWNER &&
        binding.members.find(
          (member) =>
            member === ALL_USERS_USER_EMAIL ||
            member === `user:${currentUser.value.email}`
        )
      ) {
        return true;
      }
      if (
        binding.role === PresetRoleType.PROJECT_EXPORTER &&
        binding.members.find(
          (member) =>
            member === ALL_USERS_USER_EMAIL ||
            member === `user:${currentUser.value.email}`
        )
      ) {
        if (binding.parsedExpr?.expr) {
          const conditionExpr = convertFromExpr(binding.parsedExpr.expr);
          if (
            conditionExpr.databaseResources &&
            conditionExpr.databaseResources.length > 0
          ) {
            const hasDatabaseField =
              conditionExpr.databaseResources.find(
                (item) => item.databaseName === database.name
              ) !== undefined;
            if (hasDatabaseField) {
              return true;
            }
          } else {
            return true;
          }
        } else {
          return true;
        }
      }
    });
    // If one of the binding is true, then the user is allowed to export the database.
    if (iamPolicyCheckResult.includes(true)) {
      return true;
    }

    return false;
  };

  return {
    isProjectOwnerOrDeveloper,
    isProjectOwnerOrDeveloperOrViewer,
    allowToChangeDatabaseOfProject,
    allowToQueryDatabaseV1,
    allowToExportDatabaseV1,
  };
};

import { DeploymentSchedule, Environment } from "../types";

export const generateDefaultSchedule = (environmentList: Environment[]) => {
  const schedule: DeploymentSchedule = {
    deployments: [],
  };
  environmentList.forEach((env) => {
    schedule.deployments.push({
      spec: {
        selector: {
          matchExpressions: [
            {
              key: "environment",
              operator: "In",
              values: [env.name],
            },
          ],
        },
      },
    });
  });
  return schedule;
};

/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { ExportFormat, exportFormatFromJSON, exportFormatToJSON, exportFormatToNumber } from "./common";
import { PreUpdateBackupDetail } from "./plan_check_run";

export const protobufPackage = "bytebase.store";

/** TaskDatabaseCreatePayload is the task payload for creating databases. */
export interface TaskDatabaseCreatePayload {
  /** common fields */
  skipped: boolean;
  skippedReason: string;
  specId: string;
  projectId: number;
  databaseName: string;
  tableName: string;
  sheetId: number;
  characterSet: string;
  collation: string;
  environmentId: string;
  labels: string;
}

/** TaskDatabaseDataUpdatePayload is the task payload for database data update (DML). */
export interface TaskDatabaseUpdatePayload {
  /** common fields */
  skipped: boolean;
  skippedReason: string;
  specId: string;
  schemaVersion: string;
  sheetId: number;
  preUpdateBackupDetail:
    | PreUpdateBackupDetail
    | undefined;
  /** flags is used for ghost sync */
  flags: { [key: string]: string };
}

export interface TaskDatabaseUpdatePayload_FlagsEntry {
  key: string;
  value: string;
}

/** TaskDatabaseDataExportPayload is the task payload for database data export. */
export interface TaskDatabaseDataExportPayload {
  /** common fields */
  specId: string;
  sheetId: number;
  password: string;
  format: ExportFormat;
}

function createBaseTaskDatabaseCreatePayload(): TaskDatabaseCreatePayload {
  return {
    skipped: false,
    skippedReason: "",
    specId: "",
    projectId: 0,
    databaseName: "",
    tableName: "",
    sheetId: 0,
    characterSet: "",
    collation: "",
    environmentId: "",
    labels: "",
  };
}

export const TaskDatabaseCreatePayload = {
  encode(message: TaskDatabaseCreatePayload, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.skipped === true) {
      writer.uint32(8).bool(message.skipped);
    }
    if (message.skippedReason !== "") {
      writer.uint32(18).string(message.skippedReason);
    }
    if (message.specId !== "") {
      writer.uint32(26).string(message.specId);
    }
    if (message.projectId !== 0) {
      writer.uint32(32).int32(message.projectId);
    }
    if (message.databaseName !== "") {
      writer.uint32(42).string(message.databaseName);
    }
    if (message.tableName !== "") {
      writer.uint32(50).string(message.tableName);
    }
    if (message.sheetId !== 0) {
      writer.uint32(56).int32(message.sheetId);
    }
    if (message.characterSet !== "") {
      writer.uint32(66).string(message.characterSet);
    }
    if (message.collation !== "") {
      writer.uint32(74).string(message.collation);
    }
    if (message.environmentId !== "") {
      writer.uint32(82).string(message.environmentId);
    }
    if (message.labels !== "") {
      writer.uint32(90).string(message.labels);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TaskDatabaseCreatePayload {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTaskDatabaseCreatePayload();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.skipped = reader.bool();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.skippedReason = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.specId = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.projectId = reader.int32();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.databaseName = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.tableName = reader.string();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.sheetId = reader.int32();
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.characterSet = reader.string();
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          message.collation = reader.string();
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.environmentId = reader.string();
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.labels = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TaskDatabaseCreatePayload {
    return {
      skipped: isSet(object.skipped) ? globalThis.Boolean(object.skipped) : false,
      skippedReason: isSet(object.skippedReason) ? globalThis.String(object.skippedReason) : "",
      specId: isSet(object.specId) ? globalThis.String(object.specId) : "",
      projectId: isSet(object.projectId) ? globalThis.Number(object.projectId) : 0,
      databaseName: isSet(object.databaseName) ? globalThis.String(object.databaseName) : "",
      tableName: isSet(object.tableName) ? globalThis.String(object.tableName) : "",
      sheetId: isSet(object.sheetId) ? globalThis.Number(object.sheetId) : 0,
      characterSet: isSet(object.characterSet) ? globalThis.String(object.characterSet) : "",
      collation: isSet(object.collation) ? globalThis.String(object.collation) : "",
      environmentId: isSet(object.environmentId) ? globalThis.String(object.environmentId) : "",
      labels: isSet(object.labels) ? globalThis.String(object.labels) : "",
    };
  },

  toJSON(message: TaskDatabaseCreatePayload): unknown {
    const obj: any = {};
    if (message.skipped === true) {
      obj.skipped = message.skipped;
    }
    if (message.skippedReason !== "") {
      obj.skippedReason = message.skippedReason;
    }
    if (message.specId !== "") {
      obj.specId = message.specId;
    }
    if (message.projectId !== 0) {
      obj.projectId = Math.round(message.projectId);
    }
    if (message.databaseName !== "") {
      obj.databaseName = message.databaseName;
    }
    if (message.tableName !== "") {
      obj.tableName = message.tableName;
    }
    if (message.sheetId !== 0) {
      obj.sheetId = Math.round(message.sheetId);
    }
    if (message.characterSet !== "") {
      obj.characterSet = message.characterSet;
    }
    if (message.collation !== "") {
      obj.collation = message.collation;
    }
    if (message.environmentId !== "") {
      obj.environmentId = message.environmentId;
    }
    if (message.labels !== "") {
      obj.labels = message.labels;
    }
    return obj;
  },

  create(base?: DeepPartial<TaskDatabaseCreatePayload>): TaskDatabaseCreatePayload {
    return TaskDatabaseCreatePayload.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<TaskDatabaseCreatePayload>): TaskDatabaseCreatePayload {
    const message = createBaseTaskDatabaseCreatePayload();
    message.skipped = object.skipped ?? false;
    message.skippedReason = object.skippedReason ?? "";
    message.specId = object.specId ?? "";
    message.projectId = object.projectId ?? 0;
    message.databaseName = object.databaseName ?? "";
    message.tableName = object.tableName ?? "";
    message.sheetId = object.sheetId ?? 0;
    message.characterSet = object.characterSet ?? "";
    message.collation = object.collation ?? "";
    message.environmentId = object.environmentId ?? "";
    message.labels = object.labels ?? "";
    return message;
  },
};

function createBaseTaskDatabaseUpdatePayload(): TaskDatabaseUpdatePayload {
  return {
    skipped: false,
    skippedReason: "",
    specId: "",
    schemaVersion: "",
    sheetId: 0,
    preUpdateBackupDetail: undefined,
    flags: {},
  };
}

export const TaskDatabaseUpdatePayload = {
  encode(message: TaskDatabaseUpdatePayload, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.skipped === true) {
      writer.uint32(8).bool(message.skipped);
    }
    if (message.skippedReason !== "") {
      writer.uint32(18).string(message.skippedReason);
    }
    if (message.specId !== "") {
      writer.uint32(26).string(message.specId);
    }
    if (message.schemaVersion !== "") {
      writer.uint32(34).string(message.schemaVersion);
    }
    if (message.sheetId !== 0) {
      writer.uint32(40).int32(message.sheetId);
    }
    if (message.preUpdateBackupDetail !== undefined) {
      PreUpdateBackupDetail.encode(message.preUpdateBackupDetail, writer.uint32(50).fork()).ldelim();
    }
    Object.entries(message.flags).forEach(([key, value]) => {
      TaskDatabaseUpdatePayload_FlagsEntry.encode({ key: key as any, value }, writer.uint32(58).fork()).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TaskDatabaseUpdatePayload {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTaskDatabaseUpdatePayload();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.skipped = reader.bool();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.skippedReason = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.specId = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.schemaVersion = reader.string();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.sheetId = reader.int32();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.preUpdateBackupDetail = PreUpdateBackupDetail.decode(reader, reader.uint32());
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          const entry7 = TaskDatabaseUpdatePayload_FlagsEntry.decode(reader, reader.uint32());
          if (entry7.value !== undefined) {
            message.flags[entry7.key] = entry7.value;
          }
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TaskDatabaseUpdatePayload {
    return {
      skipped: isSet(object.skipped) ? globalThis.Boolean(object.skipped) : false,
      skippedReason: isSet(object.skippedReason) ? globalThis.String(object.skippedReason) : "",
      specId: isSet(object.specId) ? globalThis.String(object.specId) : "",
      schemaVersion: isSet(object.schemaVersion) ? globalThis.String(object.schemaVersion) : "",
      sheetId: isSet(object.sheetId) ? globalThis.Number(object.sheetId) : 0,
      preUpdateBackupDetail: isSet(object.preUpdateBackupDetail)
        ? PreUpdateBackupDetail.fromJSON(object.preUpdateBackupDetail)
        : undefined,
      flags: isObject(object.flags)
        ? Object.entries(object.flags).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
    };
  },

  toJSON(message: TaskDatabaseUpdatePayload): unknown {
    const obj: any = {};
    if (message.skipped === true) {
      obj.skipped = message.skipped;
    }
    if (message.skippedReason !== "") {
      obj.skippedReason = message.skippedReason;
    }
    if (message.specId !== "") {
      obj.specId = message.specId;
    }
    if (message.schemaVersion !== "") {
      obj.schemaVersion = message.schemaVersion;
    }
    if (message.sheetId !== 0) {
      obj.sheetId = Math.round(message.sheetId);
    }
    if (message.preUpdateBackupDetail !== undefined) {
      obj.preUpdateBackupDetail = PreUpdateBackupDetail.toJSON(message.preUpdateBackupDetail);
    }
    if (message.flags) {
      const entries = Object.entries(message.flags);
      if (entries.length > 0) {
        obj.flags = {};
        entries.forEach(([k, v]) => {
          obj.flags[k] = v;
        });
      }
    }
    return obj;
  },

  create(base?: DeepPartial<TaskDatabaseUpdatePayload>): TaskDatabaseUpdatePayload {
    return TaskDatabaseUpdatePayload.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<TaskDatabaseUpdatePayload>): TaskDatabaseUpdatePayload {
    const message = createBaseTaskDatabaseUpdatePayload();
    message.skipped = object.skipped ?? false;
    message.skippedReason = object.skippedReason ?? "";
    message.specId = object.specId ?? "";
    message.schemaVersion = object.schemaVersion ?? "";
    message.sheetId = object.sheetId ?? 0;
    message.preUpdateBackupDetail =
      (object.preUpdateBackupDetail !== undefined && object.preUpdateBackupDetail !== null)
        ? PreUpdateBackupDetail.fromPartial(object.preUpdateBackupDetail)
        : undefined;
    message.flags = Object.entries(object.flags ?? {}).reduce<{ [key: string]: string }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = globalThis.String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseTaskDatabaseUpdatePayload_FlagsEntry(): TaskDatabaseUpdatePayload_FlagsEntry {
  return { key: "", value: "" };
}

export const TaskDatabaseUpdatePayload_FlagsEntry = {
  encode(message: TaskDatabaseUpdatePayload_FlagsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TaskDatabaseUpdatePayload_FlagsEntry {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTaskDatabaseUpdatePayload_FlagsEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.key = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.value = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TaskDatabaseUpdatePayload_FlagsEntry {
    return {
      key: isSet(object.key) ? globalThis.String(object.key) : "",
      value: isSet(object.value) ? globalThis.String(object.value) : "",
    };
  },

  toJSON(message: TaskDatabaseUpdatePayload_FlagsEntry): unknown {
    const obj: any = {};
    if (message.key !== "") {
      obj.key = message.key;
    }
    if (message.value !== "") {
      obj.value = message.value;
    }
    return obj;
  },

  create(base?: DeepPartial<TaskDatabaseUpdatePayload_FlagsEntry>): TaskDatabaseUpdatePayload_FlagsEntry {
    return TaskDatabaseUpdatePayload_FlagsEntry.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<TaskDatabaseUpdatePayload_FlagsEntry>): TaskDatabaseUpdatePayload_FlagsEntry {
    const message = createBaseTaskDatabaseUpdatePayload_FlagsEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseTaskDatabaseDataExportPayload(): TaskDatabaseDataExportPayload {
  return { specId: "", sheetId: 0, password: "", format: ExportFormat.FORMAT_UNSPECIFIED };
}

export const TaskDatabaseDataExportPayload = {
  encode(message: TaskDatabaseDataExportPayload, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.specId !== "") {
      writer.uint32(10).string(message.specId);
    }
    if (message.sheetId !== 0) {
      writer.uint32(16).int32(message.sheetId);
    }
    if (message.password !== "") {
      writer.uint32(26).string(message.password);
    }
    if (message.format !== ExportFormat.FORMAT_UNSPECIFIED) {
      writer.uint32(32).int32(exportFormatToNumber(message.format));
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TaskDatabaseDataExportPayload {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTaskDatabaseDataExportPayload();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.specId = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.sheetId = reader.int32();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.password = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.format = exportFormatFromJSON(reader.int32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TaskDatabaseDataExportPayload {
    return {
      specId: isSet(object.specId) ? globalThis.String(object.specId) : "",
      sheetId: isSet(object.sheetId) ? globalThis.Number(object.sheetId) : 0,
      password: isSet(object.password) ? globalThis.String(object.password) : "",
      format: isSet(object.format) ? exportFormatFromJSON(object.format) : ExportFormat.FORMAT_UNSPECIFIED,
    };
  },

  toJSON(message: TaskDatabaseDataExportPayload): unknown {
    const obj: any = {};
    if (message.specId !== "") {
      obj.specId = message.specId;
    }
    if (message.sheetId !== 0) {
      obj.sheetId = Math.round(message.sheetId);
    }
    if (message.password !== "") {
      obj.password = message.password;
    }
    if (message.format !== ExportFormat.FORMAT_UNSPECIFIED) {
      obj.format = exportFormatToJSON(message.format);
    }
    return obj;
  },

  create(base?: DeepPartial<TaskDatabaseDataExportPayload>): TaskDatabaseDataExportPayload {
    return TaskDatabaseDataExportPayload.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<TaskDatabaseDataExportPayload>): TaskDatabaseDataExportPayload {
    const message = createBaseTaskDatabaseDataExportPayload();
    message.specId = object.specId ?? "";
    message.sheetId = object.sheetId ?? 0;
    message.password = object.password ?? "";
    message.format = object.format ?? ExportFormat.FORMAT_UNSPECIFIED;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

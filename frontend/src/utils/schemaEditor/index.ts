import type { ComposedDatabase } from "@/types";
import { Engine } from "@/types/proto-es/v1/common_pb";

export * from "./filter";

export const engineSupportsSchemaEditor = (engine: Engine) => {
  if ([Engine.MYSQL, Engine.TIDB].includes(engine)) {
    return true;
  }
  if ([Engine.POSTGRES].includes(engine)) {
    return true;
  }
  return false;
};

export const allowUsingSchemaEditor = (
  databaseList: ComposedDatabase[]
): boolean => {
  return databaseList.every((db) => {
    return engineSupportsSchemaEditor(db.instanceResource.engine);
  });
};

export const getDataTypeSuggestionList = (engine: Engine = Engine.MYSQL) => {
  if (engine === Engine.MYSQL || engine === Engine.TIDB) {
    return [
      "bigint",
      "binary",
      "bit",
      "blob",
      "boolean",
      // char is equivalent to char(1).
      "char",
      "char(255)",
      "date",
      "datetime",
      "decimal",
      "double",
      "enum",
      "float",
      "geometry",
      "geometrycollection",
      "int",
      "json",
      "linestring",
      "longblob",
      "longtext",
      "mediumblob",
      "mediumint",
      "mediumtext",
      "multilinestring",
      "multipoint",
      "multipolygon",
      "point",
      "polygon",
      "smallint",
      "text",
      "time",
      "timestamp",
      "tinublob",
      "tinyint",
      "tinytext",
      // Unlike Postgres, MySQL does not support varchar with no length specified.
      "varchar(255)",
      "year",
    ];
  } else if (engine === Engine.POSTGRES) {
    return [
      "bigint",
      "bigserial",
      "bit",
      "bool",
      "boolean",
      "box",
      "bytea",
      "char",
      "char(255)",
      "cidr",
      "circle",
      "date",
      "decimal",
      "double precision",
      "float4",
      "float8",
      "inet",
      "int2",
      "int4",
      "int8",
      "integer",
      "interval",
      "json",
      "jsonb",
      "line",
      "lseg",
      "macaddr",
      "money",
      "numeric",
      "path",
      "point",
      "polygon",
      "real",
      "serial",
      "serial2",
      "serial4",
      "serial8",
      "smallint",
      "smallserial",
      "text",
      "time",
      "timestamp",
      "timestamptz",
      "timetz",
      "tsquery",
      "tsvector",
      "txid_snapshot",
      "uuid",
      "varbit",
      "varchar",
      "varchar(255)",
      "xml",
    ];
  }

  return [];
};

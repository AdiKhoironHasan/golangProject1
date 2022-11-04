/*
 Navicat Premium Data Transfer

 Source Server         : PostgreSQL
 Source Server Type    : PostgreSQL
 Source Server Version : 140005
 Source Host           : localhost:5432
 Source Catalog        : postgres
 Source Schema         : kampus

 Target Server Type    : PostgreSQL
 Target Server Version : 140005
 File Encoding         : 65001

 Date: 04/11/2022 18:42:51
*/


-- ----------------------------
-- Sequence structure for dosen_alamats_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "kampus"."dosen_alamats_id_seq";
CREATE SEQUENCE "kampus"."dosen_alamats_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "kampus"."dosen_alamats_id_seq" OWNER TO "eronman";

-- ----------------------------
-- Sequence structure for dosens_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "kampus"."dosens_id_seq";
CREATE SEQUENCE "kampus"."dosens_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "kampus"."dosens_id_seq" OWNER TO "eronman";

-- ----------------------------
-- Sequence structure for mahasiswa_alamats_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "kampus"."mahasiswa_alamats_id_seq";
CREATE SEQUENCE "kampus"."mahasiswa_alamats_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "kampus"."mahasiswa_alamats_id_seq" OWNER TO "eronman";

-- ----------------------------
-- Sequence structure for mahasiswas_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "kampus"."mahasiswas_id_seq";
CREATE SEQUENCE "kampus"."mahasiswas_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "kampus"."mahasiswas_id_seq" OWNER TO "eronman";

-- ----------------------------
-- Sequence structure for mata_kuliah_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "kampus"."mata_kuliah_id_seq";
CREATE SEQUENCE "kampus"."mata_kuliah_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "kampus"."mata_kuliah_id_seq" OWNER TO "eronman";

-- ----------------------------
-- Table structure for access_token
-- ----------------------------
DROP TABLE IF EXISTS "kampus"."access_token";
CREATE TABLE "kampus"."access_token" (
  "id" int8 NOT NULL,
  "id_user" int4,
  "token" varchar(255) COLLATE "pg_catalog"."default",
  "expired_at" timestamp(6),
  "created_at" timestamp(6)
)
;
ALTER TABLE "kampus"."access_token" OWNER TO "eronman";

-- ----------------------------
-- Records of access_token
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for dosen_alamats
-- ----------------------------
DROP TABLE IF EXISTS "kampus"."dosen_alamats";
CREATE TABLE "kampus"."dosen_alamats" (
  "id" int4 NOT NULL DEFAULT nextval('"kampus".dosen_alamats_id_seq'::regclass),
  "jalan" varchar(50) COLLATE "pg_catalog"."default",
  "no_rumah" varchar(4) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6),
  "id_dosens" int4 NOT NULL
)
;
ALTER TABLE "kampus"."dosen_alamats" OWNER TO "eronman";

-- ----------------------------
-- Records of dosen_alamats
-- ----------------------------
BEGIN;
INSERT INTO "kampus"."dosen_alamats" VALUES (1, 'jl. jawa', '23', '2022-03-21 16:49:56.263145', NULL, NULL, 1);
INSERT INTO "kampus"."dosen_alamats" VALUES (2, 'jl. merak', '26', '2022-03-21 16:49:56.263145', NULL, NULL, 1);
INSERT INTO "kampus"."dosen_alamats" VALUES (3, 'jl. jawa', '23', '2022-04-11 10:55:30.314923', NULL, NULL, 2);
INSERT INTO "kampus"."dosen_alamats" VALUES (4, 'jl. merak', '26', '2022-04-11 10:55:30.314923', NULL, NULL, 2);
INSERT INTO "kampus"."dosen_alamats" VALUES (5, 'foo ', '234', '2022-04-11 10:55:43.194442', NULL, NULL, 1);
INSERT INTO "kampus"."dosen_alamats" VALUES (7, 'jl. jawa', '23', '2022-10-07 00:15:34.135173', NULL, NULL, 4);
INSERT INTO "kampus"."dosen_alamats" VALUES (8, 'jl. merak', '26', '2022-10-07 00:15:34.135173', NULL, NULL, 4);
INSERT INTO "kampus"."dosen_alamats" VALUES (9, 'foo ', '234', '2022-10-07 00:17:11.177392', NULL, NULL, 1);
INSERT INTO "kampus"."dosen_alamats" VALUES (10, 'foo ', '234', '2022-10-07 00:17:14.492763', NULL, NULL, 1);
COMMIT;

-- ----------------------------
-- Table structure for dosens
-- ----------------------------
DROP TABLE IF EXISTS "kampus"."dosens";
CREATE TABLE "kampus"."dosens" (
  "id" int4 NOT NULL DEFAULT nextval('"kampus".dosens_id_seq'::regclass),
  "nama" varchar(50) COLLATE "pg_catalog"."default",
  "nidn" char(8) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "kampus"."dosens" OWNER TO "eronman";

-- ----------------------------
-- Records of dosens
-- ----------------------------
BEGIN;
INSERT INTO "kampus"."dosens" VALUES (1, 'Andesita', '968769  ', '2022-03-21 16:49:56.263145', NULL, NULL);
INSERT INTO "kampus"."dosens" VALUES (4, 'Andesita', '968769  ', '2022-10-07 00:15:34.135173', NULL, NULL);
INSERT INTO "kampus"."dosens" VALUES (2, 'Wahid', '968769  ', '2022-04-11 10:55:30.314923', '2022-10-07 00:17:05.212603', NULL);
COMMIT;

-- ----------------------------
-- Table structure for mahasiswa_alamats
-- ----------------------------
DROP TABLE IF EXISTS "kampus"."mahasiswa_alamats";
CREATE TABLE "kampus"."mahasiswa_alamats" (
  "id" int4 NOT NULL DEFAULT nextval('"kampus".mahasiswa_alamats_id_seq'::regclass),
  "jalan" varchar(50) COLLATE "pg_catalog"."default",
  "no_rumah" varchar(4) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6),
  "id_mahasiswas" int4 NOT NULL
)
;
ALTER TABLE "kampus"."mahasiswa_alamats" OWNER TO "eronman";

-- ----------------------------
-- Records of mahasiswa_alamats
-- ----------------------------
BEGIN;
INSERT INTO "kampus"."mahasiswa_alamats" VALUES (1, 'jalan 1 test', '1', '2021-11-25 13:07:09.682682', NULL, NULL, 1);
INSERT INTO "kampus"."mahasiswa_alamats" VALUES (2, 'jalan 2 test', '2', '2021-11-25 13:07:09.682682', NULL, NULL, 1);
INSERT INTO "kampus"."mahasiswa_alamats" VALUES (3, 'foo', '234', '2021-11-26 06:43:29.223026', NULL, NULL, 1);
INSERT INTO "kampus"."mahasiswa_alamats" VALUES (4, 'foo', '23', '2021-11-26 06:44:32.6392', NULL, NULL, 1);
INSERT INTO "kampus"."mahasiswa_alamats" VALUES (5, 'A. Yani', '45', '2021-11-26 06:53:29.560583', NULL, NULL, 1);
INSERT INTO "kampus"."mahasiswa_alamats" VALUES (30, 'foo model baru 5', '234', '2022-04-11 10:54:12.414586', NULL, NULL, 1);
INSERT INTO "kampus"."mahasiswa_alamats" VALUES (31, 'jalan 1 ulfaatest', '1', '2022-04-11 10:54:24.012198', NULL, NULL, 2);
INSERT INTO "kampus"."mahasiswa_alamats" VALUES (32, 'jalan 2 ulfaatest', '2', '2022-04-11 10:54:24.012198', NULL, NULL, 2);
INSERT INTO "kampus"."mahasiswa_alamats" VALUES (33, 'jalan 1 ulfaatest', '1', '2022-04-11 10:54:31.96782', NULL, NULL, 3);
INSERT INTO "kampus"."mahasiswa_alamats" VALUES (34, 'jalan 2 ulfaatest', '2', '2022-04-11 10:54:31.96782', NULL, NULL, 3);
INSERT INTO "kampus"."mahasiswa_alamats" VALUES (36, 'jalan 1 ulfaatest', '1', '2022-10-07 00:13:51.102405', NULL, NULL, 5);
INSERT INTO "kampus"."mahasiswa_alamats" VALUES (37, 'jalan 2 ulfaatest', '2', '2022-10-07 00:13:51.102405', NULL, NULL, 5);
INSERT INTO "kampus"."mahasiswa_alamats" VALUES (38, 'foo model baru 5', '234', '2022-10-07 00:13:55.208121', NULL, NULL, 1);
INSERT INTO "kampus"."mahasiswa_alamats" VALUES (39, 'jalan 1 ulfaatest', '1', '2022-10-07 00:16:33.994262', NULL, NULL, 6);
INSERT INTO "kampus"."mahasiswa_alamats" VALUES (40, 'jalan 2 ulfaatest', '2', '2022-10-07 00:16:33.994262', NULL, NULL, 6);
COMMIT;

-- ----------------------------
-- Table structure for mahasiswas
-- ----------------------------
DROP TABLE IF EXISTS "kampus"."mahasiswas";
CREATE TABLE "kampus"."mahasiswas" (
  "id" int4 NOT NULL DEFAULT nextval('"kampus".mahasiswas_id_seq'::regclass),
  "nama" varchar(50) COLLATE "pg_catalog"."default",
  "nim" char(8) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "kampus"."mahasiswas" OWNER TO "eronman";

-- ----------------------------
-- Records of mahasiswas
-- ----------------------------
BEGIN;
INSERT INTO "kampus"."mahasiswas" VALUES (1, 'Eko', '24402441', '2021-11-25 13:07:09.682682', '2022-04-11 10:27:28.538819', NULL);
INSERT INTO "kampus"."mahasiswas" VALUES (5, 'Eron', '2422131 ', '2022-10-07 00:13:51.102405', NULL, NULL);
INSERT INTO "kampus"."mahasiswas" VALUES (2, 'Ekoz', '2422131 ', '2022-04-11 10:54:24.012198', '2022-10-07 00:15:11.273393', NULL);
INSERT INTO "kampus"."mahasiswas" VALUES (3, 'Eronmanz', '2422131 ', '2022-04-11 10:54:31.96782', '2022-10-07 00:15:52.910743', NULL);
INSERT INTO "kampus"."mahasiswas" VALUES (6, 'Eron', '2422131 ', '2022-10-07 00:16:33.994262', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for mata_kuliah
-- ----------------------------
DROP TABLE IF EXISTS "kampus"."mata_kuliah";
CREATE TABLE "kampus"."mata_kuliah" (
  "id" int4 NOT NULL DEFAULT nextval('"kampus".mata_kuliah_id_seq'::regclass),
  "id_dosen" int4 NOT NULL,
  "nama" varchar(100) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6),
  "update_at" timestamp(6),
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "kampus"."mata_kuliah" OWNER TO "eronman";

-- ----------------------------
-- Records of mata_kuliah
-- ----------------------------
BEGIN;
INSERT INTO "kampus"."mata_kuliah" VALUES (1, 2, 'matematika', '2022-10-07 14:13:27.838897', NULL, NULL);
INSERT INTO "kampus"."mata_kuliah" VALUES (2, 2, 'matematika', '2022-10-07 14:14:08.304501', NULL, NULL);
INSERT INTO "kampus"."mata_kuliah" VALUES (3, 2, 'matematika', '2022-10-07 14:15:02.559825', NULL, NULL);
INSERT INTO "kampus"."mata_kuliah" VALUES (4, 2, 'matematika', '2022-10-07 14:20:24.859609', NULL, NULL);
INSERT INTO "kampus"."mata_kuliah" VALUES (5, 2, 'matematika', '2022-10-09 16:09:32.701282', NULL, NULL);
INSERT INTO "kampus"."mata_kuliah" VALUES (6, 2, 'matematika', '2022-10-09 16:09:57.380406', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "kampus"."users";
CREATE TABLE "kampus"."users" (
  "id" int4 NOT NULL,
  "email" varchar(100) COLLATE "pg_catalog"."default",
  "password" varchar(50) COLLATE "pg_catalog"."default",
  "name" varchar(100) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "kampus"."users" OWNER TO "eronman";

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO "kampus"."users" VALUES (1, 'eronman@gmail.com', 'Er0nM@n123', 'Adi Khoiron Hasan', '2022-11-02 23:08:15', NULL, NULL);
COMMIT;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "kampus"."dosen_alamats_id_seq"
OWNED BY "kampus"."dosen_alamats"."id";
SELECT setval('"kampus"."dosen_alamats_id_seq"', 11, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "kampus"."dosens_id_seq"
OWNED BY "kampus"."dosens"."id";
SELECT setval('"kampus"."dosens_id_seq"', 5, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "kampus"."mahasiswa_alamats_id_seq"
OWNED BY "kampus"."mahasiswa_alamats"."id";
SELECT setval('"kampus"."mahasiswa_alamats_id_seq"', 41, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "kampus"."mahasiswas_id_seq"
OWNED BY "kampus"."mahasiswas"."id";
SELECT setval('"kampus"."mahasiswas_id_seq"', 7, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "kampus"."mata_kuliah_id_seq"
OWNED BY "kampus"."mata_kuliah"."id";
SELECT setval('"kampus"."mata_kuliah_id_seq"', 7, true);

-- ----------------------------
-- Primary Key structure for table access_token
-- ----------------------------
ALTER TABLE "kampus"."access_token" ADD CONSTRAINT "access_token_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table dosens
-- ----------------------------
ALTER TABLE "kampus"."dosens" ADD CONSTRAINT "dosens_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table mahasiswas
-- ----------------------------
ALTER TABLE "kampus"."mahasiswas" ADD CONSTRAINT "mahasiswas_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table mata_kuliah
-- ----------------------------
ALTER TABLE "kampus"."mata_kuliah" ADD CONSTRAINT "mata_kuliah_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "kampus"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table dosen_alamats
-- ----------------------------
ALTER TABLE "kampus"."dosen_alamats" ADD CONSTRAINT "id_dosens_fk" FOREIGN KEY ("id_dosens") REFERENCES "kampus"."dosens" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Keys structure for table mahasiswa_alamats
-- ----------------------------
ALTER TABLE "kampus"."mahasiswa_alamats" ADD CONSTRAINT "id_mahasiswas_fk" FOREIGN KEY ("id_mahasiswas") REFERENCES "kampus"."mahasiswas" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

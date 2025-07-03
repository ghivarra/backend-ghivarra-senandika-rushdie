-- Adminer 4.8.4 PostgreSQL 15.13 (Debian 15.13-0+deb12u1) dump

DROP TABLE IF EXISTS "cart";
DROP SEQUENCE IF EXISTS cart_id_seq1;
CREATE SEQUENCE cart_id_seq1 INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."cart" (
    "id" bigint DEFAULT nextval('cart_id_seq1') NOT NULL,
    "user_id" bigint NOT NULL,
    "product_id" bigint NOT NULL,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    "deleted_at" timestamp,
    "quantity" integer DEFAULT '1' NOT NULL,
    CONSTRAINT "cart_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "cart_deleted_at" ON "public"."cart" USING btree ("deleted_at");

CREATE INDEX "cart_product_id" ON "public"."cart" USING btree ("product_id");

CREATE INDEX "cart_user_id" ON "public"."cart" USING btree ("user_id");


DROP TABLE IF EXISTS "invoice";
DROP SEQUENCE IF EXISTS invoice_id_seq;
CREATE SEQUENCE invoice_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."invoice" (
    "id" bigint DEFAULT nextval('invoice_id_seq') NOT NULL,
    "price" integer NOT NULL,
    "details" text NOT NULL,
    "user_id" bigint NOT NULL,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    "deleted_at" timestamp,
    CONSTRAINT "invoice_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "invoice_deleted_at" ON "public"."invoice" USING btree ("deleted_at");

CREATE INDEX "invoice_user_id" ON "public"."invoice" USING btree ("user_id");


DROP TABLE IF EXISTS "invoice_order_list";
DROP SEQUENCE IF EXISTS invoice_order_list_id_seq;
CREATE SEQUENCE invoice_order_list_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."invoice_order_list" (
    "id" bigint DEFAULT nextval('invoice_order_list_id_seq') NOT NULL,
    "invoice_id" bigint NOT NULL,
    "order_id" bigint NOT NULL,
    CONSTRAINT "invoice_order_list_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "invoice_order_list_invoice_id" ON "public"."invoice_order_list" USING btree ("invoice_id");

CREATE INDEX "invoice_order_list_order_id" ON "public"."invoice_order_list" USING btree ("order_id");


DROP TABLE IF EXISTS "order";
DROP SEQUENCE IF EXISTS order_id_seq;
CREATE SEQUENCE order_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."order" (
    "id" bigint DEFAULT nextval('order_id_seq') NOT NULL,
    "buyer_id" bigint NOT NULL,
    "seller_id" bigint NOT NULL,
    "product_id" bigint NOT NULL,
    "price" integer NOT NULL,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    "deleted_at" timestamp,
    "quantity" bigint NOT NULL,
    CONSTRAINT "order_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "order_buyer_id" ON "public"."order" USING btree ("buyer_id");

CREATE INDEX "order_deleted_at" ON "public"."order" USING btree ("deleted_at");

CREATE INDEX "order_product_id" ON "public"."order" USING btree ("product_id");

CREATE INDEX "order_seller_id" ON "public"."order" USING btree ("seller_id");


DROP TABLE IF EXISTS "product";
DROP SEQUENCE IF EXISTS product_id_seq;
CREATE SEQUENCE product_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."product" (
    "id" bigint DEFAULT nextval('product_id_seq') NOT NULL,
    "name" character varying(200) NOT NULL,
    "description" text,
    "photo" character varying(200) NOT NULL,
    "price" integer NOT NULL,
    "stock" integer NOT NULL,
    "slug" character varying(200) NOT NULL,
    "user_id" bigint NOT NULL,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    "deleted_at" timestamp,
    CONSTRAINT "product_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "product_slug" UNIQUE ("slug")
) WITH (oids = false);

CREATE INDEX "product_deleted_at" ON "public"."product" USING btree ("deleted_at");

CREATE INDEX "product_name" ON "public"."product" USING btree ("name");

CREATE INDEX "product_user_id" ON "public"."product" USING btree ("user_id");


DROP TABLE IF EXISTS "user";
DROP SEQUENCE IF EXISTS user_id_seq;
CREATE SEQUENCE user_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."user" (
    "id" bigint DEFAULT nextval('user_id_seq') NOT NULL,
    "username" character varying(100) NOT NULL,
    "password" character varying(200) NOT NULL,
    "name" character varying(100) NOT NULL,
    "email" character varying(100) NOT NULL,
    "user_role_id" bigint NOT NULL,
    "is_active" smallint DEFAULT '1' NOT NULL,
    "created_at" timestamp DEFAULT now(),
    "updated_at" timestamp DEFAULT now(),
    "deleted_at" timestamp,
    CONSTRAINT "user_email" UNIQUE ("email"),
    CONSTRAINT "user_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "user_username" UNIQUE ("username")
) WITH (oids = false);

CREATE INDEX "user_deleted_at" ON "public"."user" USING btree ("deleted_at");

CREATE INDEX "user_is_active" ON "public"."user" USING btree ("is_active");

CREATE INDEX "user_user_role_id" ON "public"."user" USING btree ("user_role_id");


DROP TABLE IF EXISTS "user_module";
DROP SEQUENCE IF EXISTS user_module_id_seq1;
CREATE SEQUENCE user_module_id_seq1 INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."user_module" (
    "id" bigint DEFAULT nextval('user_module_id_seq1') NOT NULL,
    "name" character varying(200) NOT NULL,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    "deleted_at" timestamp,
    CONSTRAINT "user_module_name" UNIQUE ("name"),
    CONSTRAINT "user_module_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "user_module_deleted_at" ON "public"."user_module" USING btree ("deleted_at");

INSERT INTO "user_module" ("id", "name", "created_at", "updated_at", "deleted_at") VALUES
(1,	'/admin/product/create',	'2025-07-03 02:19:45.004789',	'2025-07-03 02:19:45.004789',	NULL),
(2,	'/admin/product/update',	'2025-07-03 02:20:15.163835',	'2025-07-03 02:20:15.163835',	NULL),
(3,	'/admin/product/delete',	'2025-07-03 02:20:23.641968',	'2025-07-03 02:20:23.641968',	NULL),
(5,	'/cart/add-product',	'2025-07-03 04:39:23.448502',	'2025-07-03 04:39:23.448502',	NULL),
(4,	'/cart/',	'2025-07-03 04:39:18.112237',	'2025-07-03 04:39:18.112237',	NULL),
(6,	'/cart/buy',	'2025-07-03 05:18:12.893902',	'2025-07-03 05:18:12.893902',	NULL),
(7,	'/order/',	'2025-07-03 05:18:24.548542',	'2025-07-03 05:18:24.548542',	NULL);

DROP TABLE IF EXISTS "user_role";
DROP SEQUENCE IF EXISTS user_role_id_seq1;
CREATE SEQUENCE user_role_id_seq1 INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."user_role" (
    "id" bigint DEFAULT nextval('user_role_id_seq1') NOT NULL,
    "name" character varying(50) NOT NULL,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    "deleted_at" timestamp,
    CONSTRAINT "user_role_name" UNIQUE ("name"),
    CONSTRAINT "user_role_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "user_role_deleted_at" ON "public"."user_role" USING btree ("deleted_at");

INSERT INTO "user_role" ("id", "name", "created_at", "updated_at", "deleted_at") VALUES
(1,	'Merchant',	'2025-07-03 00:50:22.998435',	'2025-07-03 00:50:22.998435',	NULL),
(2,	'Customer',	'2025-07-03 00:50:22.998435',	'2025-07-03 00:50:22.998435',	NULL);

DROP TABLE IF EXISTS "user_role_module_list";
DROP SEQUENCE IF EXISTS user_role_module_list_id_seq;
CREATE SEQUENCE user_role_module_list_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."user_role_module_list" (
    "id" bigint DEFAULT nextval('user_role_module_list_id_seq') NOT NULL,
    "user_role_id" bigint NOT NULL,
    "user_module_id" bigint NOT NULL,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    "deleted_at" timestamp,
    CONSTRAINT "user_role_module_list_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "user_role_module_list_deleted_at" ON "public"."user_role_module_list" USING btree ("deleted_at");

CREATE INDEX "user_role_module_list_user_module_id" ON "public"."user_role_module_list" USING btree ("user_module_id");

CREATE INDEX "user_role_module_list_user_role_id" ON "public"."user_role_module_list" USING btree ("user_role_id");

INSERT INTO "user_role_module_list" ("id", "user_role_id", "user_module_id", "created_at", "updated_at", "deleted_at") VALUES
(1,	1,	1,	'2025-07-03 02:41:44.193298',	'2025-07-03 02:41:44.193298',	NULL),
(2,	1,	2,	'2025-07-03 02:41:46.481891',	'2025-07-03 02:41:46.481891',	NULL),
(3,	1,	3,	'2025-07-03 02:41:48.327253',	'2025-07-03 02:41:48.327253',	NULL),
(4,	2,	4,	'2025-07-03 04:39:41.504892',	'2025-07-03 04:39:41.504892',	NULL),
(5,	2,	5,	'2025-07-03 04:39:44.129005',	'2025-07-03 04:39:44.129005',	NULL),
(6,	2,	6,	'2025-07-03 05:18:36.700769',	'2025-07-03 05:18:36.700769',	NULL),
(7,	1,	7,	'2025-07-03 05:18:39.599075',	'2025-07-03 05:18:39.599075',	NULL);

ALTER TABLE ONLY "public"."cart" ADD CONSTRAINT "cart_product_id_fkey" FOREIGN KEY (product_id) REFERENCES product(id) ON UPDATE CASCADE ON DELETE CASCADE NOT DEFERRABLE;
ALTER TABLE ONLY "public"."cart" ADD CONSTRAINT "cart_user_id_fkey" FOREIGN KEY (user_id) REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE CASCADE NOT DEFERRABLE;

ALTER TABLE ONLY "public"."order" ADD CONSTRAINT "order_buyer_id_fkey" FOREIGN KEY (buyer_id) REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE RESTRICT NOT DEFERRABLE;
ALTER TABLE ONLY "public"."order" ADD CONSTRAINT "order_product_id_fkey" FOREIGN KEY (product_id) REFERENCES product(id) ON UPDATE CASCADE ON DELETE RESTRICT NOT DEFERRABLE;
ALTER TABLE ONLY "public"."order" ADD CONSTRAINT "order_seller_id_fkey" FOREIGN KEY (seller_id) REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE RESTRICT NOT DEFERRABLE;

ALTER TABLE ONLY "public"."product" ADD CONSTRAINT "product_user_id_fkey" FOREIGN KEY (user_id) REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE RESTRICT NOT DEFERRABLE;

ALTER TABLE ONLY "public"."user" ADD CONSTRAINT "user_user_role_id_fkey" FOREIGN KEY (user_role_id) REFERENCES user_role(id) ON UPDATE CASCADE ON DELETE RESTRICT NOT DEFERRABLE;

-- 2025-07-03 07:19:28.459921+07

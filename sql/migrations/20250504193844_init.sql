-- Create "authors" table
CREATE TABLE "public"."authors" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(255) NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_authors_deleted_at" to table: "authors"
CREATE INDEX "idx_authors_deleted_at" ON "public"."authors" ("deleted_at");
-- Create "books" table
CREATE TABLE "public"."books" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "isbn" character varying(13) NOT NULL,
  "title" character varying(255) NOT NULL,
  "author_id" uuid NOT NULL,
  "copies" bigint NOT NULL DEFAULT 1,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_authors_books" FOREIGN KEY ("author_id") REFERENCES "public"."authors" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_books_deleted_at" to table: "books"
CREATE INDEX "idx_books_deleted_at" ON "public"."books" ("deleted_at");
-- Create index "idx_books_isbn" to table: "books"
CREATE UNIQUE INDEX "idx_books_isbn" ON "public"."books" ("isbn");
-- Create "borrowers" table
CREATE TABLE "public"."borrowers" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(255) NOT NULL,
  "email" character varying(255) NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_borrowers_deleted_at" to table: "borrowers"
CREATE INDEX "idx_borrowers_deleted_at" ON "public"."borrowers" ("deleted_at");
-- Create index "idx_borrowers_email" to table: "borrowers"
CREATE UNIQUE INDEX "idx_borrowers_email" ON "public"."borrowers" ("email");
-- Create "loans" table
CREATE TABLE "public"."loans" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "book_id" uuid NOT NULL,
  "borrower_id" uuid NOT NULL,
  "due_date" timestamptz NOT NULL,
  "returned_at" timestamptz NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_borrowers_loans" FOREIGN KEY ("borrower_id") REFERENCES "public"."borrowers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_loans_book" FOREIGN KEY ("book_id") REFERENCES "public"."books" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_loans_deleted_at" to table: "loans"
CREATE INDEX "idx_loans_deleted_at" ON "public"."loans" ("deleted_at");

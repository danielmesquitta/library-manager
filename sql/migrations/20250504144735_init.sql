-- Create "authors" table
CREATE TABLE "public"."authors" (
  "id" bigserial NOT NULL,
  "name" character varying(255) NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "books" table
CREATE TABLE "public"."books" (
  "id" bigserial NOT NULL,
  "isbn" character varying(13) NOT NULL,
  "title" character varying(255) NOT NULL,
  "author_id" bigint NULL,
  "copies" bigint NULL DEFAULT 1,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_authors_books" FOREIGN KEY ("author_id") REFERENCES "public"."authors" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_books_isbn" to table: "books"
CREATE UNIQUE INDEX "idx_books_isbn" ON "public"."books" ("isbn");
-- Create "borrowers" table
CREATE TABLE "public"."borrowers" (
  "id" bigserial NOT NULL,
  "name" character varying(255) NOT NULL,
  "email" character varying(255) NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_borrowers_email" to table: "borrowers"
CREATE UNIQUE INDEX "idx_borrowers_email" ON "public"."borrowers" ("email");
-- Create "loans" table
CREATE TABLE "public"."loans" (
  "id" bigserial NOT NULL,
  "book_id" bigint NULL,
  "borrower_id" bigint NULL,
  "due_date" timestamptz NULL,
  "returned_at" timestamptz NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_borrowers_loans" FOREIGN KEY ("borrower_id") REFERENCES "public"."borrowers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_loans_book" FOREIGN KEY ("book_id") REFERENCES "public"."books" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);

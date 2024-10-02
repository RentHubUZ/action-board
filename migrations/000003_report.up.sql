CREATE TABLE IF NOT EXISTS "reports" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "property_id" uuid NOT NULL,
  "user_id" UUID NOT NULL,
  "email" varchar NOT NULL,
  "problem" varchar NOT NULL,
  "description" text,
  "created_at" TIMESTAMP DEFAULT (NOW())
);
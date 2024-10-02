CREATE TABLE IF NOT EXISTS "notifications" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "user_id" UUID NOT NULL,
  "content" text,
  "is_read" bool,
  "created_at" TIMESTAMP DEFAULT (NOW())
);
CREATE TABLE "requests" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "user_id" UUID NOT NULL,
  "email" varchar[] NOT NULL,
  "roommate_count" int,
  "phone_number" varchar
);
CREATE TABLE "reviews" (
  "id" UUID PRIMARY KEY DEFAULT (gen_random_uuid()),
  "user_id" UUID NOT NULL, --buni tekshirish kerak emas ya'ni bu id tokendan keladi
  "property_id" UUID NOT NULL,  --bu columni properties tabledan tekshirish kerak
  "rating" INT,
  "comment" TEXT,
  "created_at" TIMESTAMP DEFAULT NOW(),
  "updated_at" TIMESTAMP DEFAULT NOW()
);
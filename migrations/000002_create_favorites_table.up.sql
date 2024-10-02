CREATE Table if NOT EXISTS "favorites" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" UUID not null, --buni tekshirish kerak emas ya'ni bu id tokendan keladi
  "property_id" UUID not null, --bu columni properties tabledan tekshirish kerak
  "created_at" TIMESTAMP DEFAULT NOW()
);
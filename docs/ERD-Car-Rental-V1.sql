CREATE TABLE "customers" (
  "id" int PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "nik" varchar(20) UNIQUE NOT NULL,
  "phone_number" varchar(15) NOT NULL
);

CREATE TABLE "cars" (
  "id" int PRIMARY KEY,
  "brand" varchar(50) NOT NULL,
  "type" "enum(SUV,Sedan,MPV,Hatchback,Crossover,Wagon,Pickup,Coupe,Van,Convertible)" NOT NULL,
  "transmission" "enum(Manual,Automatic)" NOT NULL,
  "plate_number" varchar(20) UNIQUE NOT NULL,
  "price_per_day" int NOT NULL,
  "available" boolean NOT NULL
);

CREATE TABLE "bookings" (
  "id" int PRIMARY KEY,
  "customer_id" int NOT NULL,
  "car_id" int NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "total_price" int NOT NULL,
  "status" "enum(Pending,Confirmed,Completed,Cancelled)"
);

ALTER TABLE "bookings" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id") DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "bookings" ADD FOREIGN KEY ("car_id") REFERENCES "cars" ("id") DEFERRABLE INITIALLY IMMEDIATE;

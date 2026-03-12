# car-rental-v1
Car Rental API built with Golang and PostgreSQL implementing CRUD operations for Customers, Cars, and Bookings.

This system is designed to manage the car rental process in a structured way, including the management of customer data (Customers), vehicle data (Cars), and rental transactions (Bookings).

A. Folder and File Structure
    1. Config: Contains application configurations related to database connection:
        .env → stores environment variables for the database.
        config.go → loads the application configuration.

    2. Controllers: Handles client requests and generates API responses:
        customersController.go, carsController.go, bookingsController.go

    3. Database: Manages database connections and table migrations:
        create_customers.sql, create_cars.sql, create_bookings.sql, database.go

    4. Docs: Contains system documentation and database diagrams: 
        ERD-Car-Rental-V1.png, RED-Car-Rental-V1.sql

    5. Models: Defines data structures representing database tables:
        customers.go, cars.go, bookings.go

    6. Repositories: Handles database operations such as create, read, update, and delete (CRUD).

    7. Routes: Defines API endpoints and connects them to the corresponding controllers.
    
    8. Utils: Contains helper functions that support various processes in the application.

    9. main.go: The entry point of the program that runs the API server and initializes configuration and database connections.

    10. go.mod: The dependency management file in Golang that defines the project module, Go version, and all external packages used, ensuring the CAR RENTAL V1 API runs correctly.

B. API Overview
    1. Customers API: Manages customer data.
        POST /customers → add a new customer.
        GET /customers → retrieve all customers.
        PUT /customers/:id → update customer data by ID.
        DELETE /customers/:id → delete customer by ID.

    2. Cars API: Manages vehicle data available for rent.
        POST /cars → add a new car.
        GET /cars → retrieve all cars.
        PUT /cars/:id → update car data by ID.
        DELETE /cars/:id → delete a car by ID.

    3. Bookings API: Manages rental transaction data.
        POST /bookings → create a new booking.
        GET /bookings → retrieve all bookings.
        PUT /bookings/:id → update booking by ID.
        DELETE /bookings/:id → delete booking by ID.

C. Database Structure and Relationships

    1. Customers: Stores customer data
        id → unique identifier.
        name → customer name.
        nik → unique identification number.
        phone_number → contact number.
            A customer can have multiple bookings, so the relationship to the bookings table is one-to-many.

    2. Cars: Stores available vehicle data
        id → unique identifier.
        brand → car brand.
        type → car type (SUV, Sedan, MPV, Hatchback, Crossover, Wagon, Pickup, Coupe, Van, Convertible).
        transmission → transmission type (Manual or Automatic).
        plate_number → unique license plate number.
        price_per_day → rental price per day.
        available → indicates whether the car is available.
            A car can be rented multiple times by different customers at different dates, so the relationship to bookings is also one-to-many.

    3. Bookings: Records all rental transactions
        id → unique identifier.
        customer_id → foreign key to customers.id.
        car_id → foreign key to cars.id.
        start_date and end_date → rental period.
        total_price → rental cost.
        status → booking status (Pending, Confirmed, Completed, Cancelled).
            This table acts as a junction table between customers and cars, ensuring each transaction is clearly recorded: who rented, which car, rental period, cost, and status.

Relationships and Data Flow
    One Customer can create one or more bookings → one-to-many.
    One Car can be rented multiple times by different customers at different periods → one-to-many.
    The Bookings table serves as the central connector, recording all transactions, linking customer and car data, and tracking rental periods and booking status.
    The overall relationship pattern: Customers → Bookings ← Cars, where Bookings is the core that manages the entire rental workflow.


Note: SQL/ERD and PostgreSQL differ because ERD is conceptual, representing entities, attributes, and relationships abstractly, whereas PostgreSQL is a real implementation requiring valid syntax, data types, auto-increment, foreign keys, and constraints. Therefore, ERD structures must be adjusted when applied to a live PostgreSQL database.


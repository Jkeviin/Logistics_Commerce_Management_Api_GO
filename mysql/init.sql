CREATE DATABASE IF NOT EXISTS `group_10_db`
/*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */
/*!80016 DEFAULT ENCRYPTION='N' */
;

USE `group_10_db`;

-- MySQL dump 10.13  Distrib 8.0.25, for Linux (x86_64)
--
-- Host: localhost    Database: group_10_db
-- ------------------------------------------------------
-- Server version   8.0.29-0ubuntu0.20.04.3
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */
;

/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */
;

/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */
;

/*!50503 SET NAMES utf8 */
;

/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */
;

/*!40103 SET TIME_ZONE='+00:00' */
;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */
;

/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */
;

/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */
;

/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */
;

-- UTF8 en todas las tablas ajustes completo, para que agarre tildes y todo:
ALTER DATABASE group_10_db CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

CREATE TABLE products_type (
    id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL
);

CREATE TABLE buyer (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_card_number INT NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    CONSTRAINT uq_buyer_id_card_number UNIQUE (id_card_number)
);

CREATE TABLE locality (
    id VARCHAR(255) PRIMARY KEY,
    locality_name VARCHAR(255) NOT NULL,
    province_name VARCHAR(255) NOT NULL,
    country_name VARCHAR(255) NOT NULL
);

CREATE TABLE warehouse (
    id INT AUTO_INCREMENT PRIMARY KEY,
    warehouse_code VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    telephone VARCHAR(255) NOT NULL,
    minimum_capacity INT NOT NULL,
    minimum_temperature DECIMAL(19, 2) NOT NULL,
    locality_id VARCHAR(255) NOT NULL,
    UNIQUE KEY uq_warehouse_code (warehouse_code),
    FOREIGN KEY (locality_id) REFERENCES locality(id) -- Clave foránea añadida
);

CREATE TABLE seller (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cid VARCHAR(255) NOT NULL,
    company_name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    telephone VARCHAR(255) NOT NULL,
    locality_id VARCHAR(255) NOT NULL,
    UNIQUE KEY uq_seller_cid (cid),
    FOREIGN KEY (locality_id) REFERENCES locality(id)
);

CREATE TABLE carry (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cid VARCHAR(255) NOT NULL,
    company_name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    telephone VARCHAR(255) NOT NULL,
    locality_id VARCHAR(255) NOT NULL,
    UNIQUE KEY uq_carry_cid (cid),
    FOREIGN KEY (locality_id) REFERENCES locality(id) -- Clave foránea añadida
);

CREATE TABLE employee (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_card_number INT NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    warehouse_id INT NOT NULL,
    UNIQUE KEY uq_card_number (id_card_number),
    FOREIGN KEY (warehouse_id) REFERENCES warehouse(id) -- Clave foránea añadida
);

CREATE TABLE product (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_code VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    width DECIMAL(19, 2) NOT NULL,
    height DECIMAL(19, 2) NOT NULL,
    length DECIMAL(19, 2) NOT NULL,
    netweight DECIMAL(19, 2) NOT NULL,
    expiration_rate DECIMAL(19, 2) NOT NULL,
    recommended_freezing_temperature DECIMAL(19, 2) NOT NULL,
    freezing_rate DECIMAL(19, 2) NOT NULL,
    product_type_id INT NOT NULL,
    seller_id INT NOT NULL,
    expiration_date DATETIME NOT NULL,
    UNIQUE KEY uq_product_code (product_code),
    FOREIGN KEY (product_type_id) REFERENCES products_type(id),
    FOREIGN KEY (seller_id) REFERENCES seller(id)
);

CREATE TABLE section (
    id INT AUTO_INCREMENT PRIMARY KEY,
    section_number VARCHAR(255) NOT NULL,
    current_temperature DECIMAL(19, 2) NOT NULL,
    minimum_temperature DECIMAL(19, 2) NOT NULL,
    current_capacity INT NOT NULL,
    minimum_capacity INT NOT NULL,
    maximum_capacity INT NOT NULL,
    warehouse_id INT NOT NULL,
    product_type_id INT NOT NULL,
    UNIQUE KEY uq_section_number (section_number),
    FOREIGN KEY (product_type_id) REFERENCES products_type(id),
    FOREIGN KEY (warehouse_id) REFERENCES warehouse(id)
);

-- Corrección de tipos DATE/DATETIME
CREATE TABLE product_batch (
    id INT AUTO_INCREMENT PRIMARY KEY,
    batch_number VARCHAR(255) NOT NULL,
    current_quantity INT NOT NULL,
    current_temperature DECIMAL(19, 2) NOT NULL,
    due_date DATE NOT NULL,
    initial_quantity INT NOT NULL,
    manufacturing_date DATE NOT NULL,
    manufacturing_hour INT NOT NULL,
    minimum_temperature DECIMAL(19, 2) NOT NULL,
    product_id INT NOT NULL,
    section_id INT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product(id),
    FOREIGN KEY (section_id) REFERENCES section(id)
);

CREATE TABLE product_record (
    id INT AUTO_INCREMENT PRIMARY KEY,
    last_update_date DATETIME NOT NULL,
    purchase_price DECIMAL(19, 2) NOT NULL,
    sale_price DECIMAL(19, 2) NOT NULL,
    product_id INT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product(id)
);

CREATE TABLE inbound_order (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_date DATE NOT NULL,
    order_number VARCHAR(255) NOT NULL,
    temperature DECIMAL(19, 2) NOT NULL,
    employee_id INT NOT NULL,
    product_batch_id INT NOT NULL,
    warehouse_id INT NOT NULL,
    UNIQUE KEY uq_order_number (order_number),
    FOREIGN KEY (employee_id) REFERENCES employee(id),
    FOREIGN KEY (product_batch_id) REFERENCES product_batch(id),
    FOREIGN KEY (warehouse_id) REFERENCES warehouse(id)
);

CREATE TABLE purchase_order (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_number VARCHAR(255) NOT NULL,
    order_date DATE NOT NULL,
    tracking_code VARCHAR(255) NOT NULL,
    buyer_id INT NOT NULL,
    product_record_id INT NOT NULL,
    UNIQUE KEY uq_purchase_order_number (order_number),
    FOREIGN KEY (buyer_id) REFERENCES buyer(id),
    FOREIGN KEY (product_record_id) REFERENCES product_record(id)
);


-- Tabla: products_type
INSERT INTO
    products_type (description)
VALUES
    ('Lácteos');

INSERT INTO
    products_type (description)
VALUES
    ('Carnes');

INSERT INTO
    products_type (description)
VALUES
    ('Vegetales');

-- Tabla: buyer
INSERT INTO
    buyer (id_card_number, first_name, last_name)
VALUES
    (12345678, 'Juan', 'Pérez');

INSERT INTO
    buyer (id_card_number, first_name, last_name)
VALUES
    (87654321, 'Ana', 'García');

INSERT INTO
    buyer (id_card_number, first_name, last_name)
VALUES
    (11223344, 'Luis', 'Martínez');

-- Tabla: locality
INSERT INTO
    locality (id, locality_name, province_name, country_name)
VALUES
    ('B1643', 'Palermo', 'Buenos Aires', 'Argentina');

INSERT INTO
    locality (id, locality_name, province_name, country_name)
VALUES
    ('B1644','Recoleta', 'Buenos Aires', 'Argentina');

INSERT INTO
    locality (id, locality_name, province_name, country_name)
VALUES
    ('B1645', 'Belgrano', 'Buenos Aires', 'Argentina');

-- Tabla: warehouse
INSERT INTO
    warehouse (
        warehouse_code,
        address,
        telephone,
        minimum_capacity,
        minimum_temperature,
        locality_id
    )
VALUES
    (
        'WH001',
        'Calle 123',
        '1111-1111',
        100,
        -5.0,
        (
            SELECT
                id
            FROM
                locality
            WHERE
                locality_name = 'Palermo'
            LIMIT
                1
        )
    );

INSERT INTO
    warehouse (
        warehouse_code,
        address,
        telephone,
        minimum_capacity,
        minimum_temperature,
        locality_id
    )
VALUES
    (
        'WH002',
        'Avenida 456',
        '2222-2222',
        200,
        -10.0,
        (
            SELECT
                id
            FROM
                locality
            WHERE
                locality_name = 'Recoleta'
            LIMIT
                1
        )
    );

INSERT INTO
    warehouse (
        warehouse_code,
        address,
        telephone,
        minimum_capacity,
        minimum_temperature,
        locality_id
    )
VALUES
    (
        'WH003',
        'Boulevard 789',
        '3333-3333',
        150,
        -8.0,
        (
            SELECT
                id
            FROM
                locality
            WHERE
                locality_name = 'Belgrano'
            LIMIT
                1
        )
    );

-- Tabla: seller
INSERT INTO
    seller (
        cid,
        company_name,
        address,
        telephone,
        locality_id
    )
VALUES
    (
        'C001',
        'Lácteos S.A.',
        'Calle 1',
        '4444-4444',
        (
            SELECT
                id
            FROM
                locality
            WHERE
                locality_name = 'Palermo'
            LIMIT
                1
        )
    );

INSERT INTO
    seller (
        cid,
        company_name,
        address,
        telephone,
        locality_id
    )
VALUES
    (
        'C002',
        'Carnes S.R.L.',
        'Calle 2',
        '5555-5555',
        (
            SELECT
                id
            FROM
                locality
            WHERE
                locality_name = 'Recoleta'
            LIMIT
                1
        )
    );

INSERT INTO
    seller (
        cid,
        company_name,
        address,
        telephone,
        locality_id
    )
VALUES
    (
        'C003',
        'Verde Vida',
        'Calle 3',
        '6666-6666',
        (
            SELECT
                id
            FROM
                locality
            WHERE
                locality_name = 'Belgrano'
            LIMIT
                1
        )
    );

-- Tabla: carry
INSERT INTO
    carry (
        cid,
        company_name,
        address,
        telephone,
        locality_id
    )
VALUES
    (
        'T001',
        'Transporte Uno',
        'Ruta 1',
        '7777-7777',
        (
            SELECT
                id
            FROM
                locality
            WHERE
                locality_name = 'Palermo'
            LIMIT
                1
        )
    );

INSERT INTO
    carry (
        cid,
        company_name,
        address,
        telephone,
        locality_id
    )
VALUES
    (
        'T002',
        'Transporte Dos',
        'Ruta 2',
        '8888-8888',
        (
            SELECT
                id
            FROM
                locality
            WHERE
                locality_name = 'Recoleta'
            LIMIT
                1
        )
    );

INSERT INTO
    carry (
        cid,
        company_name,
        address,
        telephone,
        locality_id
    )
VALUES
    (
        'T003',
        'Transporte Tres',
        'Ruta 3',
        '9999-9999',
        (
            SELECT
                id
            FROM
                locality
            WHERE
                locality_name = 'Belgrano'
            LIMIT
                1
        )
    );

-- Tabla: employee
INSERT INTO
    employee (
        id_card_number,
        first_name,
        last_name,
        warehouse_id
    )
VALUES
    (
        22334455,
        'Carlos',
        'Ruiz',
        (
            SELECT
                id
            FROM
                warehouse
            WHERE
                warehouse_code = 'WH001'
            LIMIT
                1
        )
    );

INSERT INTO
    employee (
        id_card_number,
        first_name,
        last_name,
        warehouse_id
    )
VALUES
    (
        33445566,
        'María',
        'López',
        (
            SELECT
                id
            FROM
                warehouse
            WHERE
                warehouse_code = 'WH002'
            LIMIT
                1
        )
    );

INSERT INTO
    employee (
        id_card_number,
        first_name,
        last_name,
        warehouse_id
    )
VALUES
    (
        44556677,
        'Pedro',
        'Gómez',
        (
            SELECT
                id
            FROM
                warehouse
            WHERE
                warehouse_code = 'WH003'
            LIMIT
                1
        )
    );

-- Tabla: product
INSERT INTO
    product (
        product_code,
        description,
        width,
        height,
        length,
        netweight,
        expiration_rate,
        recommended_freezing_temperature,
        freezing_rate,
        product_type_id,
        seller_id,
        expiration_date
    )
VALUES
    (
        'P001',
        'Leche Entera',
        10.0,
        20.0,
        30.0,
        1.0,
        0.5,
        -4.0,
        0.2,
        (
            SELECT
                id
            FROM
                products_type
            WHERE
                description = 'Lácteos'
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                seller
            WHERE
                company_name = 'Lácteos S.A.'
            LIMIT
                1
        ), '2025-01-01 00:00:00'
    );

INSERT INTO
    product (
        product_code,
        description,
        width,
        height,
        length,
        netweight,
        expiration_rate,
        recommended_freezing_temperature,
        freezing_rate,
        product_type_id,
        seller_id,
        expiration_date
    )
VALUES
    (
        'P002',
        'Carne Vacuna',
        15.0,
        25.0,
        35.0,
        2.0,
        0.7,
        -8.0,
        0.3,
        (
            SELECT
                id
            FROM
                products_type
            WHERE
                description = 'Carnes'
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                seller
            WHERE
                company_name = 'Carnes S.R.L.'
            LIMIT
                1
        ), '2025-02-01 00:00:00'
    );

INSERT INTO
    product (
        product_code,
        description,
        width,
        height,
        length,
        netweight,
        expiration_rate,
        recommended_freezing_temperature,
        freezing_rate,
        product_type_id,
        seller_id,
        expiration_date
    )
VALUES
    (
        'P003',
        'Zanahoria',
        5.0,
        10.0,
        15.0,
        0.5,
        0.2,
        2.0,
        0.1,
        (
            SELECT
                id
            FROM
                products_type
            WHERE
                description = 'Vegetales'
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                seller
            WHERE
                company_name = 'Verde Vida'
            LIMIT
                1
        ), '2025-03-01 00:00:00'
    );

-- Tabla: section
INSERT INTO
    section (
        section_number,
        current_temperature,
        minimum_temperature,
        current_capacity,
        minimum_capacity,
        maximum_capacity,
        warehouse_id,
        product_type_id
    )
VALUES
    (
        'S001',
        -4.0,
        -5.0,
        50,
        20,
        100,
        (
            SELECT
                id
            FROM
                warehouse
            WHERE
                warehouse_code = 'WH001'
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                products_type
            WHERE
                description = 'Lácteos'
            LIMIT
                1
        )
    );

INSERT INTO
    section (
        section_number,
        current_temperature,
        minimum_temperature,
        current_capacity,
        minimum_capacity,
        maximum_capacity,
        warehouse_id,
        product_type_id
    )
VALUES
    (
        'S002',
        -8.0,
        -10.0,
        80,
        30,
        150,
        (
            SELECT
                id
            FROM
                warehouse
            WHERE
                warehouse_code = 'WH002'
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                products_type
            WHERE
                description = 'Carnes'
            LIMIT
                1
        )
    );

INSERT INTO
    section (
        section_number,
        current_temperature,
        minimum_temperature,
        current_capacity,
        minimum_capacity,
        maximum_capacity,
        warehouse_id,
        product_type_id
    )
VALUES
    (
        'S003',
        2.0,
        0.0,
        60,
        25,
        120,
        (
            SELECT
                id
            FROM
                warehouse
            WHERE
                warehouse_code = 'WH003'
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                products_type
            WHERE
                description = 'Vegetales'
            LIMIT
                1
        )
    );

-- Tabla: product_batch
INSERT INTO
    product_batch (
        batch_number,
        current_quantity,
        current_temperature,
        due_date,
        initial_quantity,
        manufacturing_date,
        manufacturing_hour,
        minimum_temperature,
        product_id,
        section_id
    )
VALUES
    (
        'B001',
        30,
        -4.0,
        '2025-01-10',
        50,
        '2024-12-01',
        8,
        -5.0,
        (
            SELECT
                id
            FROM
                product
            WHERE
                product_code = 'P001'
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                section
            WHERE
                section_number = 'S001'
            LIMIT
                1
        )
    );

INSERT INTO
    product_batch (
        batch_number,
        current_quantity,
        current_temperature,
        due_date,
        initial_quantity,
        manufacturing_date,
        manufacturing_hour,
        minimum_temperature,
        product_id,
        section_id
    )
VALUES
    (
        'B002',
        40,
        -8.0,
        '2025-02-10',
        60,
        '2024-12-15',
        9,
        -10.0,
        (
            SELECT
                id
            FROM
                product
            WHERE
                product_code = 'P002'
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                section
            WHERE
                section_number = 'S002'
            LIMIT
                1
        )
    );

INSERT INTO
    product_batch (
        batch_number,
        current_quantity,
        current_temperature,
        due_date,
        initial_quantity,
        manufacturing_date,
        manufacturing_hour,
        minimum_temperature,
        product_id,
        section_id
    )
VALUES
    (
        'B003',
        20,
        2.0,
        '2025-03-10 00:00:00',
        30,
        '2024-12-20',
        '10:00:00',
        0.0,
        (
            SELECT
                id
            FROM
                product
            WHERE
                product_code = 'P003'
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                section
            WHERE
                section_number = 'S003'
            LIMIT
                1
        )
    );

-- Tabla: product_record
INSERT INTO
    product_record (
        last_update_date,
        purchase_price,
        sale_price,
        product_id
    )
VALUES
    (
        '2024-12-05 00:00:00',
        100.00,
        120.00,
        (
            SELECT
                id
            FROM
                product
            WHERE
                product_code = 'P001'
            LIMIT
                1
        )
    );

INSERT INTO
    product_record (
        last_update_date,
        purchase_price,
        sale_price,
        product_id
    )
VALUES
    (
        '2024-12-10 00:00:00',
        200.00,
        250.00,
        (
            SELECT
                id
            FROM
                product
            WHERE
                product_code = 'P002'
            LIMIT
                1
        )
    );

INSERT INTO
    product_record (
        last_update_date,
        purchase_price,
        sale_price,
        product_id
    )
VALUES
    (
        '2024-12-15 00:00:00',
        50.00,
        70.00,
        (
            SELECT
                id
            FROM
                product
            WHERE
                product_code = 'P003'
            LIMIT
                1
        )
    );

-- Tabla: inbound_order
INSERT INTO
    inbound_order (
        order_date,
        order_number,
        temperature,
        employee_id,
        product_batch_id,
        warehouse_id
    )
VALUES
    (
        '2024-12-06 00:00:00',
        'IO001',
        -4.0,
        (
            SELECT
                id
            FROM
                employee
            WHERE
                id_card_number = 22334455
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                product_batch
            WHERE
                batch_number = 'B001'
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                warehouse
            WHERE
                warehouse_code = 'WH001'
            LIMIT
                1
        )
    );

INSERT INTO
    inbound_order (
        order_date,
        order_number,
        temperature,
        employee_id,
        product_batch_id,
        warehouse_id
    )
VALUES
    (
        '2024-12-11 00:00:00',
        'IO002',
        -8.0,
        (
            SELECT
                id
            FROM
                employee
            WHERE
                id_card_number = 33445566
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                product_batch
            WHERE
                batch_number = 'B002'
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                warehouse
            WHERE
                warehouse_code = 'WH002'
            LIMIT
                1
        )
    );

INSERT INTO
    inbound_order (
        order_date,
        order_number,
        temperature,
        employee_id,
        product_batch_id,
        warehouse_id
    )
VALUES
    (
        '2024-12-16 00:00:00',
        'IO003',
        2.0,
        (
            SELECT
                id
            FROM
                employee
            WHERE
                id_card_number = 44556677
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                product_batch
            WHERE
                batch_number = 'B003'
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                warehouse
            WHERE
                warehouse_code = 'WH003'
            LIMIT
                1
        )
    );

-- Tabla: purchase_order
INSERT INTO
    purchase_order (
        order_number,
        order_date,
        tracking_code,
        buyer_id,
        product_record_id
    )
VALUES
    (
        'PO001',
        '2024-12-07',
        'TRK001',
        (
            SELECT
                id
            FROM
                buyer
            WHERE
                id_card_number = 12345678
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                product_record
            WHERE
                product_id =(
                    SELECT
                        id
                    FROM
                        product
                    WHERE
                        product_code = 'P001'
                    LIMIT
                        1
                )
            LIMIT
                1
        )
    );

INSERT INTO
    purchase_order (
        order_number,
        order_date,
        tracking_code,
        buyer_id,
        product_record_id
    )
VALUES
    (
        'PO002',
        '2024-12-12',
        'TRK002',
        (
            SELECT
                id
            FROM
                buyer
            WHERE
                id_card_number = 87654321
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                product_record
            WHERE
                product_id =(
                    SELECT
                        id
                    FROM
                        product
                    WHERE
                        product_code = 'P002'
                    LIMIT
                        1
                )
            LIMIT
                1
        )
    );

INSERT INTO
    purchase_order (
        order_number,
        order_date,
        tracking_code,
        buyer_id,
        product_record_id
    )
VALUES
    (
        'PO003',
        '2024-12-17',
        'TRK003',
        (
            SELECT
                id
            FROM
                buyer
            WHERE
                id_card_number = 11223344
            LIMIT
                1
        ), (
            SELECT
                id
            FROM
                product_record
            WHERE
                product_id =(
                    SELECT
                        id
                    FROM
                        product
                    WHERE
                        product_code = 'P003'
                    LIMIT
                        1
                )
            LIMIT
                1
        )
    );
-- Create 'users' table
CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    phone_number VARCHAR(15) UNIQUE NOT NULL,
    address TEXT,
    pin VARCHAR(100) NOT NULL,
    balance BIGINT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert example users
INSERT INTO users (first_name, last_name, phone_number, address, pin, balance) 
VALUES 
('Guntur', 'Saputro', '0811255501', 'Jl. Kebon Sirih No. 1', '$2a$10$eBzOiGJZbux6BHKLE3vVOOBZaKziDRxRZ.VMaNEopYUN3Vjso8VtO', 500000),  -- PIN: 123456 (hashed with bcrypt)
('Tom', 'Araya', '0811255502', 'Jl. Diponegoro No. 215', '$2a$10$eBzOiGJZbux6BHKLE3vVOOBZaKziDRxRZ.VMaNEopYUN3Vjso8VtO', 300000);  -- PIN: 123456 (hashed with bcrypt)

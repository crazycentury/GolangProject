-- Create 'transactions' table
CREATE TABLE transactions (
    transaction_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(user_id),
    target_user_id UUID,
    amount BIGINT,
    type VARCHAR(10), -- CREDIT/DEBIT
    remarks TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert example transactions
INSERT INTO transactions (user_id, target_user_id, amount, type, remarks) 
VALUES 
('bc1c823e-b0fb-4b20-88c0-dff25e283252', 'b7342e8e-e8e7-4a5d-873e-b1b1bfcdeddb', 30000, 'DEBIT', 'Hadiah Ultah'),
('bc1c823e-b0fb-4b20-88c0-dff25e283252', NULL, 500000, 'CREDIT', 'Top up balance');
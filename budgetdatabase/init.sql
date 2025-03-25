-- Table des clients et de leur budget total
CREATE TABLE budgets (
    budgets_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL, 
    category TEXT NOT NULL, 
    budgets_limit DECIMAL(10, 2) NOT NULL, 
    current_amount DECIMAL(10, 2) DEFAULT 0
);



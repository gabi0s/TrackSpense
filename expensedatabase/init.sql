CREATE TABLE expenses (
    expense_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    budget_id UUID NOT NULL,
    budget_categorie VARCHAR(100) NOT NULL,
    price DECIMAL(10,2) NOT NULL
);

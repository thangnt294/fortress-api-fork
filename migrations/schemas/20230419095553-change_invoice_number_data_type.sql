-- +migrate Up
ALTER TABLE invoices ALTER COLUMN discount TYPE DECIMAL USING discount::DECIMAL;
ALTER TABLE invoices ALTER COLUMN tax TYPE DECIMAL USING tax::DECIMAL;
ALTER TABLE invoices ALTER COLUMN sub_total TYPE DECIMAL USING sub_total::DECIMAL;

-- +migrate Down
ALTER TABLE invoices ALTER COLUMN discount TYPE INTEGER USING discount::INTEGER;
ALTER TABLE invoices ALTER COLUMN tax TYPE INTEGER USING tax::INTEGER;
ALTER TABLE invoices ALTER COLUMN sub_total TYPE INTEGER USING sub_total::INTEGER;
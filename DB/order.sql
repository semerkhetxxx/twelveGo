-- 建立 order 表（如果不存在）
CREATE TABLE IF NOT EXISTS "order" (
    -- 訂單的唯一標識符
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- 下單用戶的唯一標識符
    user_id UUID NOT NULL,
    -- 訂單日期
    order_date TIMESTAMP NOT NULL,
    -- 稅金金額
    tax_amount DECIMAL(10, 2) NOT NULL,
    -- 運費金額
    shipping_fee DECIMAL(10, 2) NOT NULL,
    -- 訂單總額
    total_amount DECIMAL(10, 2) NOT NULL,
    -- 預計送達日期
    estimated_delivery_date DATE,
    -- 訂單狀態（待付款、已付款、已發貨、已完成等）
    status VARCHAR(20) NOT NULL
);

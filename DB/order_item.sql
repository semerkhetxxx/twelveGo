-- 建立 order_item 表（如果不存在）
CREATE TABLE IF NOT EXISTS order_item (
    -- 訂單明細的唯一標識符
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- 所屬訂單的唯一標識符
    order_id UUID NOT NULL,
    -- 商品的唯一標識符
    product_id UUID NOT NULL,
    -- 商品連結
    product_link VARCHAR(255) NOT NULL,
    -- 商品名稱
    product_name VARCHAR(100) NOT NULL,
    -- 商品副標/簡易說明
    product_subtitle VARCHAR(255),
    -- 商品單價
    unit_price DECIMAL(10, 2) NOT NULL,
    -- 購買數量
    quantity INTEGER NOT NULL
);
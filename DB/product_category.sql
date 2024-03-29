
-- 建立 product_category 表（如果不存在）
CREATE TABLE IF NOT EXISTS product_category (
    -- 商品分類的唯一標識符
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- 商品分類名稱
    name TEXT NOT NULL,
    -- 商品分類的描述
    description TEXT
);
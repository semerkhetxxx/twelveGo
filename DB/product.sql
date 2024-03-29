-- 建立 product 表（如果不存在）
CREATE TABLE IF NOT EXISTS product (
    -- 商品的唯一標識符
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- 商品名稱
    name VARCHAR(100) NOT NULL,
    -- 商品價格
    price DECIMAL(10, 2) NOT NULL,
    -- 商品庫存數量
    stock_quantity INTEGER NOT NULL,
    -- 商品分類的唯一標識符
    category_id UUID NOT NULL,
    -- 商品描述
    description TEXT,
    -- 圖片/影片連結 1
    media_link_1 VARCHAR(255),
    -- 圖片/影片連結 2
    media_link_2 VARCHAR(255),
    -- 圖片/影片連結 3
    media_link_3 VARCHAR(255),
    -- 圖片/影片連結 4
    media_link_4 VARCHAR(255),
    -- 圖片/影片連結 5
    media_link_5 VARCHAR(255),
    -- 圖片/影片連結 6
    media_link_6 VARCHAR(255)
);

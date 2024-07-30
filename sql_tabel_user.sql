CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);


INSERT INTO users (name, email, phone, created_at, updated_at)
VALUES 
('Agus Santoso', 'agus.santoso@example.com', '081234567890', NOW(), NOW()),
('Budi Setiawan', 'budi.setiawan@example.com', '081234567891', NOW(), NOW()),
('Citra Dewi', 'citra.dewi@example.com', '081234567892', NOW(), NOW()),
('Dewi Lestari', 'dewi.lestari@example.com', '081234567893', NOW(), NOW()),
('Eko Prasetyo', 'eko.prasetyo@example.com', '081234567894', NOW(), NOW()),
('Fajar Nugroho', 'fajar.nugroho@example.com', '081234567895', NOW(), NOW()),
('Gita Putri', 'gita.putri@example.com', '081234567896', NOW(), NOW()),
('Hendra Wijaya', 'hendra.wijaya@example.com', '081234567897', NOW(), NOW()),
('Intan Permata', 'intan.permata@example.com', '081234567898', NOW(), NOW()),
('Joko Susilo', 'joko.susilo@example.com', '081234567899', NOW(), NOW());

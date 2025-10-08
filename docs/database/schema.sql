DROP TABLE IF EXISTS refunds CASCADE;
DROP TABLE IF EXISTS booking_seats CASCADE;
DROP TABLE IF EXISTS bookings CASCADE;
DROP TABLE IF EXISTS schedules CASCADE;
DROP TABLE IF EXISTS seats CASCADE;
DROP TABLE IF EXISTS studios CASCADE;
DROP TABLE IF EXISTS movies CASCADE;
DROP TABLE IF EXISTS cinemas CASCADE;
DROP TABLE IF EXISTS cities CASCADE;
DROP TABLE IF EXISTS users CASCADE;

DROP TYPE IF EXISTS user_role_enum CASCADE;
DROP TYPE IF EXISTS seat_type_enum CASCADE;
DROP TYPE IF EXISTS schedule_status_enum CASCADE;
DROP TYPE IF EXISTS booking_status_enum CASCADE;
DROP TYPE IF EXISTS booking_seat_status_enum CASCADE;
DROP TYPE IF EXISTS refund_status_enum CASCADE;

CREATE TYPE user_role_enum AS ENUM ('customer', 'admin', 'staff');
CREATE TYPE seat_type_enum AS ENUM ('regular', 'vip', 'premium');
CREATE TYPE schedule_status_enum AS ENUM ('active', 'cancelled', 'completed');
CREATE TYPE booking_status_enum AS ENUM ('pending', 'paid', 'cancelled', 'refunded');
CREATE TYPE booking_seat_status_enum AS ENUM ('locked', 'sold', 'available');
CREATE TYPE refund_status_enum AS ENUM ('pending', 'approved', 'rejected', 'completed');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    role user_role_enum DEFAULT 'customer',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);

CREATE TABLE cities (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    code VARCHAR(10) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_cities_code ON cities(code);

CREATE TABLE cinemas (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    city_id UUID NOT NULL REFERENCES cities(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    address TEXT,
    phone VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_cinemas_city_id ON cinemas(city_id);

CREATE TABLE studios (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cinema_id UUID NOT NULL REFERENCES cinemas(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    total_seats INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_studios_cinema_id ON studios(cinema_id);

CREATE TABLE seats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    studio_id UUID NOT NULL REFERENCES studios(id) ON DELETE CASCADE,
    seat_number VARCHAR(10) NOT NULL,
    row VARCHAR(5) NOT NULL,
    seat_type seat_type_enum DEFAULT 'regular',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(studio_id, seat_number)
);

CREATE INDEX idx_seats_studio_id ON seats(studio_id);

CREATE TABLE movies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    duration INTEGER NOT NULL,
    genre VARCHAR(100),
    rating VARCHAR(10),
    poster_url VARCHAR(500),
    release_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_movies_title ON movies(title);
CREATE INDEX idx_movies_release_date ON movies(release_date);

CREATE TABLE schedules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    movie_id UUID NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    studio_id UUID NOT NULL REFERENCES studios(id) ON DELETE CASCADE,
    show_date DATE NOT NULL,
    show_time TIME NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    status schedule_status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(studio_id, show_date, show_time)
);

CREATE INDEX idx_schedules_movie_id ON schedules(movie_id);
CREATE INDEX idx_schedules_studio_id ON schedules(studio_id);
CREATE INDEX idx_schedules_show_date ON schedules(show_date);
CREATE INDEX idx_schedules_status ON schedules(status);

CREATE TABLE bookings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    schedule_id UUID NOT NULL REFERENCES schedules(id) ON DELETE CASCADE,
    booking_code VARCHAR(20) UNIQUE NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    status booking_status_enum DEFAULT 'pending',
    payment_method VARCHAR(50),
    paid_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_bookings_user_id ON bookings(user_id);
CREATE INDEX idx_bookings_schedule_id ON bookings(schedule_id);
CREATE INDEX idx_bookings_booking_code ON bookings(booking_code);
CREATE INDEX idx_bookings_status ON bookings(status);

CREATE TABLE booking_seats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
    seat_id UUID NOT NULL REFERENCES seats(id) ON DELETE CASCADE,
    schedule_id UUID NOT NULL REFERENCES schedules(id) ON DELETE CASCADE,
    status booking_seat_status_enum DEFAULT 'locked',
    locked_at TIMESTAMP,
    locked_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(seat_id, schedule_id)
);

CREATE INDEX idx_booking_seats_booking_id ON booking_seats(booking_id);
CREATE INDEX idx_booking_seats_seat_id ON booking_seats(seat_id);
CREATE INDEX idx_booking_seats_schedule_id ON booking_seats(schedule_id);
CREATE INDEX idx_booking_seats_status ON booking_seats(status);
CREATE INDEX idx_booking_seats_locked_at ON booking_seats(locked_at);

CREATE TABLE refunds (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
    amount DECIMAL(10,2) NOT NULL,
    reason TEXT,
    refund_method VARCHAR(50),
    status refund_status_enum DEFAULT 'pending',
    processed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_refunds_booking_id ON refunds(booking_id);
CREATE INDEX idx_refunds_status ON refunds(status);

INSERT INTO cities (id, name, code, created_at) VALUES
('550e8400-e29b-41d4-a716-446655440001', 'Jakarta', 'JKT', NOW()),
('550e8400-e29b-41d4-a716-446655440002', 'Bandung', 'BDG', NOW()),
('550e8400-e29b-41d4-a716-446655440003', 'Surabaya', 'SBY', NOW());

INSERT INTO cinemas (id, city_id, name, address, phone, created_at, updated_at) VALUES
('550e8400-e29b-41d4-a716-446655440011', '550e8400-e29b-41d4-a716-446655440001', 'Cinema XXI Grand Indonesia', 'Jl. M.H. Thamrin No. 1, Jakarta Pusat', '021-2358-1234', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440012', '550e8400-e29b-41d4-a716-446655440001', 'CGV Pacific Place', 'Jl. Jend. Sudirman Kav. 52-53, Jakarta Selatan', '021-5797-1234', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440013', '550e8400-e29b-41d4-a716-446655440002', 'Cinepolis Paris Van Java', 'Jl. Sukajadi No. 137-139, Bandung', '022-2033-1234', NOW(), NOW());

INSERT INTO studios (id, cinema_id, name, total_seats, created_at) VALUES
('550e8400-e29b-41d4-a716-446655440021', '550e8400-e29b-41d4-a716-446655440011', 'Studio 1', 150, NOW()),
('550e8400-e29b-41d4-a716-446655440022', '550e8400-e29b-41d4-a716-446655440011', 'Studio 2', 100, NOW()),
('550e8400-e29b-41d4-a716-446655440023', '550e8400-e29b-41d4-a716-446655440012', 'Premium Studio', 80, NOW()),
('550e8400-e29b-41d4-a716-446655440024', '550e8400-e29b-41d4-a716-446655440013', 'Velvet Studio', 60, NOW());

INSERT INTO movies (id, title, description, duration, genre, rating, poster_url, release_date, created_at, updated_at) VALUES
('550e8400-e29b-41d4-a716-446655440031', 'Avengers: Endgame', 'The epic conclusion to the Infinity Saga', 181, 'Action, Adventure, Sci-Fi', '13+', 'https://www.imdb.com/title/tt4154796/mediaviewer/rm2775147008/?ref_=tt_ov_i', '2025-04-26', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440032', 'Spider-Man: No Way Home', 'Peter Parker seeks Doctor Stranges help', 148, 'Action, Adventure, Fantasy', '13+', 'https://www.imdb.com/title/tt10872600/mediaviewer/rm2798542593/?ref_=tt_ov_i', '2025-12-17', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440033', 'The Batman', 'Batman ventures into Gothams underworld', 176, 'Action, Crime, Drama', '17+', 'https://www.imdb.com/title/tt1877830/mediaviewer/rm3486063105/?ref_=tt_ov_i', '2025-03-04', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440034', 'Dune: Part Two', 'Paul Atreides unites with the Fremen', 166, 'Action, Adventure, Sci-Fi', '13+', 'https://www.imdb.com/title/tt15239678/mediaviewer/rm3705359361/?ref_=tt_ov_i', '2025-03-01', NOW(), NOW());

INSERT INTO users (id, email, password, full_name, phone, role, created_at, updated_at) VALUES
('550e8400-e29b-41d4-a716-446655440041', 'admin@gmail.com', '$2a$12$8rgVucKN1ICNqYH78JxcC.iFNi5ewg6H2wWj34o7HxLg4jmSBEMBu', 'Admin Cinema', '081234567890', 'admin', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440042', 'staff@gmail.com', '$2a$12$8rgVucKN1ICNqYH78JxcC.iFNi5ewg6H2wWj34o7HxLg4jmSBEMBu', 'Staff Cinema', '081234567891', 'staff', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440043', 'customer@gmail.com', '$2a$12$8rgVucKN1ICNqYH78JxcC.iFNi5ewg6H2wWj34o7HxLg4jmSBEMBu', 'Customer Cinema', '081234567892', 'customer', NOW(), NOW());

SELECT 'cities' as table_name, COUNT(*) as row_count FROM cities
UNION ALL
SELECT 'cinemas', COUNT(*) FROM cinemas
UNION ALL
SELECT 'studios', COUNT(*) FROM studios
UNION ALL
SELECT 'movies', COUNT(*) FROM movies
UNION ALL
SELECT 'users', COUNT(*) FROM users;
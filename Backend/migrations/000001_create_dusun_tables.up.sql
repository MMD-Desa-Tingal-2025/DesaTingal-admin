CREATE TABLE dusun (
    id UUID PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    warna VARCHAR(20),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE statistik_dusun (
    id UUID PRIMARY KEY,
    dusun_id UUID NOT NULL,
    jumlah_kk INTEGER NOT NULL,
    jumlah_laki INTEGER NOT NULL,
    jumlah_perempuan INTEGER NOT NULL,
    jumlah_rumah INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_dusun FOREIGN KEY (dusun_id) REFERENCES dusun(id) ON DELETE CASCADE
);

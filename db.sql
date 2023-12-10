CREATE TYPE tipe_catatan AS ENUM (
	'masuk', 
	'keluar'
);

/* Menampung catatan setiap mobil
yang masuk/keluar
di tempat parkir. */
CREATE TABLE catatan (
	id bigserial PRIMARY KEY,
	tipe tipe_catatan NOT NULL,
	waktu timestamp DEFAULT LOCALTIMESTAMP NOT NULL,
	picc integer NOT NULL,
	nama character(32) NOT NULL
);

/* Menampung uid dan nama dari kartu picc
yang menjadi identitas mobil
yang sedang parkir. */
CREATE TABLE terparkir (
	picc integer PRIMARY KEY,
	nama character(32) NOT NULL
);

-- Menampung nilai tarif dan saat mulai berlakunya.
CREATE TABLE tarif (
	id smallserial PRIMARY KEY,
	tarif smallint NOT NULL,
	berlaku timestamp DEFAULT LOCALTIMESTAMP NOT NULL
);

/* Memasukkan baris baru ke catatan
dengan tipe catatan t, uid picc p, dan nama n. */
CREATE PROCEDURE catat(t tipe_catatan, p integer, n character(32)) 
AS $$
	INSERT INTO catatan (tipe, picc, nama) VALUES (t, p, n);
$$
LANGUAGE SQL;

-- Memunculkan error dengan pesannya.
CREATE FUNCTION galat(m character(128))
RETURNS void
AS $$
BEGIN
	RAISE EXCEPTION 'GALAT: %', m;
	RETURN;
END;
$$
LANGUAGE plpgsql;

-- Memunculkan error insert-only.
CREATE FUNCTION galat_insert_only()
RETURNS TRIGGER
AS $$
BEGIN
	CALL galat('Tabel tersebut bersifat insert-only');
	RETURN OLD;
END;
$$
LANGUAGE plpgsql;

/* Memasukkan baris baru ke catatan
dengan tipe catatan masuk,
picc dan nama bersumber dari
baris yang baru dimasukkan ke tabel terparkir. */
CREATE FUNCTION catat_masuk()
RETURNS TRIGGER
AS $$
BEGIN
	CALL catat('masuk', NEW.picc, NEW.nama);
	RETURN NEW;
END;
$$
LANGUAGE plpgsql;

/* Memasukkan baris baru ke catatan
dengan tipe catatan keluar,
picc dan nama bersumber dari
baris yang baru dihapus dari tabel terparkir. */
CREATE FUNCTION catat_keluar()
RETURNS TRIGGER
AS $$
BEGIN
	CALL catat('keluar', OLD.picc, OLD.nama);
	RETURN NEW;
END;
$$
LANGUAGE plpgsql;

-- Mencatat apabila ada mobil yang masuk.
CREATE TRIGGER masuk
AFTER INSERT ON terparkir
FOR EACH ROW
EXECUTE PROCEDURE catat_masuk();

-- Mencatat apabila ada mobil yang keluar.
CREATE TRIGGER keluar
AFTER DELETE ON terparkir
FOR EACH ROW
EXECUTE PROCEDURE catat_keluar();

-- Memastikan tabel catatan bersifat insert-only.
CREATE TRIGGER catatan_insert_only
BEFORE UPDATE OR DELETE ON catatan
FOR EACH STATEMENT
EXECUTE PROCEDURE galat_insert_only();

-- Memastikan tabel tarif bersifat insert-only.
CREATE TRIGGER tarif_insert_only
BEFORE UPDATE OR DELETE ON tarif
FOR EACH STATEMENT
EXECUTE PROCEDURE galat_insert_only();

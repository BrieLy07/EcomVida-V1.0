CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL
);

CREATE TABLE permisos (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL
);

CREATE TABLE rol_permiso (
    rol_id INT REFERENCES roles(id),
    permiso_id INT REFERENCES permisos(id),
    PRIMARY KEY (rol_id, permiso_id)
);

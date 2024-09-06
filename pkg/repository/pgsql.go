package repository

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PGRepo struct {
	mu   sync.Mutex
	pool *pgxpool.Pool
}

func New(connStr string) (*PGRepo, error) {
	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	_, err = pool.Exec(context.Background(), "CREATE TABLE genres(id serial primary key, name text not null default 'Без названия');")
	if err != nil {
		return nil, err
	}
	_, err = pool.Exec(context.Background(), "CREATE TABLE authors(id serial primary key, name text not null default 'Инкогнито');")
	if err != nil {
		return nil, err
	}
	_, err = pool.Exec(context.Background(), "CREATE TABLE books(id serial primary key, name text not null default 'Без названия', author_id integer not null references authors(id), genre_id integer not null references genres(id), price integer not null default 100);")
	if err != nil {
		return nil, err
	}
	_, err = pool.Exec(context.Background(), "INSERT INTO authors (name) values ('Ф. Достоевский'), ('М. Булгаков'),('А. Пушкин'),('Р. Брэдбери');")
	if err != nil {
		return nil, err
	}
	_, err = pool.Exec(context.Background(), "	INSERT INTO genres (name) values ('Роман'), ('Рассказ'), ('Роман в стихах'), ('Антиутопия');")
	if err != nil {
		return nil, err
	}
	_, err = pool.Exec(context.Background(), "	INSERT INTO books (name, author_id, genre_id) values ('Преступление и наказание', 1, 1);")
	if err != nil {
		return nil, err
	}
	return &PGRepo{mu: sync.Mutex{}, pool: pool}, nil
}

package main

type Storage interface {
	PersistFileData(*File) error
	GetFileData(*File) error
}

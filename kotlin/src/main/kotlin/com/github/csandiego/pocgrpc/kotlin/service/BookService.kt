package com.github.csandiego.pocgrpc.kotlin.service

import com.github.csandiego.pocgrpc.kotlin.data.Book

interface BookService {

    suspend fun list(): List<Book>

    suspend fun get(id: Int): Book

    suspend fun create(book: Book)

    suspend fun update(id: Int, book: Book)

    suspend fun delete(id: Int)
}
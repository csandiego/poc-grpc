package com.github.csandiego.pocgrpc.kotlin.grpc

import com.github.csandiego.pocgrpc.kotlin.protobuf.Book as DTO
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookCreateRequest
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookDeleteRequest
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookGetRequest
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookListRequest
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookServiceGrpc.BookServiceBlockingStub
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookUpdateRequest
import com.github.csandiego.pocgrpc.kotlin.data.Book
import com.github.csandiego.pocgrpc.kotlin.service.BookService

class GrpcBlockingBookService(private val stub: BookServiceBlockingStub) : BookService {

    override suspend fun list(): List<Book> {
        return stub.list(BookListRequest.newBuilder().build()).getBooksList().map {
            Book(it.id.toInt(), it.title, it.author)
        }
    }

    override suspend fun get(id: Int): Book {
        val request = BookGetRequest.newBuilder().setId(id.toLong()).build()
        return stub.get(request).getBook().let { Book(it.id.toInt(), it.title, it.author)}
    }

    override suspend fun create(book: Book) {
        val dto = DTO.newBuilder().setTitle(book.title).setAuthor(book.author).build()
        val request = BookCreateRequest.newBuilder().setBook(dto).build()
        stub.create(request)
    }

    override suspend fun update(id: Int, book: Book) {
        val dto = DTO.newBuilder().setTitle(book.title).setAuthor(book.author).build()
        val request = BookUpdateRequest.newBuilder().setId(id.toLong()).setBook(dto).build()
        stub.update(request)
    }

    override suspend fun delete(id: Int) {
        val request = BookDeleteRequest.newBuilder().setId(id.toLong()).build()
        stub.delete(request)
    }
}
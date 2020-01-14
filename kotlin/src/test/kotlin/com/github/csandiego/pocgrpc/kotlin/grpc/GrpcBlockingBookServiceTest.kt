package com.github.csandiego.pocgrpc.kotlin.grpc

import com.github.csandiego.pocgrpc.kotlin.data.Book
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookCreateRequest
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookCreateResponse
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookDeleteRequest
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookDeleteResponse
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookGetRequest
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookGetResponse
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookListRequest
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookListResponse
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookServiceGrpc
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookUpdateRequest
import com.github.csandiego.pocgrpc.kotlin.protobuf.BookUpdateResponse
import io.grpc.inprocess.InProcessChannelBuilder
import io.grpc.inprocess.InProcessServerBuilder
import io.grpc.stub.StreamObserver
import io.grpc.testing.GrpcCleanupRule
import kotlinx.coroutines.runBlocking
import org.junit.Before
import org.junit.Rule
import org.junit.Test
import kotlin.test.assertEquals
import com.github.csandiego.pocgrpc.kotlin.protobuf.Book as DTO

class GrpcBlockingBookServiceTest {

    private lateinit var server: TestBookServiceServer
    private lateinit var service: GrpcBlockingBookService

    @get:Rule
    val grpcCleanup = GrpcCleanupRule()

    @Before
    fun setUp() {
        server = TestBookServiceServer()

        val serverName = InProcessServerBuilder.generateName()
        grpcCleanup.register(InProcessServerBuilder.forName(serverName).directExecutor().addService(server).build().start())
        val channel = grpcCleanup.register(InProcessChannelBuilder.forName(serverName).directExecutor().build())
        val stub = BookServiceGrpc.newBlockingStub(channel)

        service = GrpcBlockingBookService(stub)
    }

    @Test
    fun list() = runBlocking {
        val dto = listOf(DTO.newBuilder().setId(1).setTitle("abc").setAuthor("xyz").build())
        server.listResponse = BookListResponse.newBuilder().addAllBooks(dto).build()
        val books = service.list()
        assertEquals(dto[0].id, books[0].id!!.toLong())
        assertEquals(dto[0].title, books[0].title)
        assertEquals(dto[0].author, books[0].author)
    }

    @Test
    fun get() = runBlocking {
        val id = 1
        val dto = DTO.newBuilder().setId(id.toLong()).setTitle("abc").setAuthor("xyz").build()
        server.getResponse = BookGetResponse.newBuilder().setBook(dto).build()
        val book = service.get(id)
        assertEquals(id.toLong(), server.getRequest.id)
        assertEquals(dto.id, book.id!!.toLong())
        assertEquals(dto.title, book.title)
        assertEquals(dto.author, book.author)
    }

    @Test
    fun create() = runBlocking {
        val book = Book(null, "abc", "xyz")
        server.createResponse = BookCreateResponse.newBuilder().build()
        service.create(book)
        assertEquals(book.title, server.createRequest.book.title)
        assertEquals(book.author, server.createRequest.book.author)
    }

    @Test
    fun update() = runBlocking {
        val id = 1
        val book = Book(null, "abc", "xyz")
        server.updateResponse = BookUpdateResponse.newBuilder().build()
        service.update(id, book)
        assertEquals(id.toLong(), server.updateRequest.id)
        assertEquals(book.title, server.updateRequest.book.title)
        assertEquals(book.author, server.updateRequest.book.author)
    }

    @Test
    fun delete() = runBlocking {
        val id = 1
        server.deleteResponse = BookDeleteResponse.newBuilder().build()
        service.delete(id)
        assertEquals(id.toLong(), server.deleteRequest.id)
    }

    class TestBookServiceServer : BookServiceGrpc.BookServiceImplBase() {

        lateinit var listRequest: BookListRequest
        lateinit var listResponse: BookListResponse
        lateinit var createRequest: BookCreateRequest
        lateinit var createResponse: BookCreateResponse
        lateinit var getRequest: BookGetRequest
        lateinit var getResponse: BookGetResponse
        lateinit var updateRequest: BookUpdateRequest
        lateinit var updateResponse: BookUpdateResponse
        lateinit var deleteRequest: BookDeleteRequest
        lateinit var deleteResponse: BookDeleteResponse

        override fun list(request: BookListRequest, responseObserver: StreamObserver<BookListResponse>) {
            listRequest = request
            with(responseObserver) {
                onNext(listResponse)
                onCompleted()
            }
        }

        override fun create(request: BookCreateRequest, responseObserver: StreamObserver<BookCreateResponse>) {
            createRequest = request
            with(responseObserver) {
                onNext(createResponse)
                onCompleted()
            }

        }

        override fun get(request: BookGetRequest, responseObserver: StreamObserver<BookGetResponse>) {
            getRequest = request
            with(responseObserver) {
                onNext(getResponse)
                onCompleted()
            }

        }

        override fun update(request: BookUpdateRequest, responseObserver: StreamObserver<BookUpdateResponse>) {
            updateRequest = request
            with(responseObserver) {
                onNext(updateResponse)
                onCompleted()
            }

        }

        override fun delete(request: BookDeleteRequest, responseObserver: StreamObserver<BookDeleteResponse>) {
            deleteRequest = request
            with(responseObserver) {
                onNext(deleteResponse)
                onCompleted()
            }

        }
    }
}
//
//  GRPCBookService.swift
//  POC GRPC SwiftKitTests
//
//  Created by Christopher San Diego on 1/15/20.
//  Copyright © 2020 Christopher San Diego. All rights reserved.
//

import XCTest
import GRPC
import NIO
import POC_GRPC_SwiftKit

class GRPCBookServiceTests: XCTestCase {
    
    private var serverGroup: EventLoopGroup!
    private var clientGroup: EventLoopGroup!
    private var server: Server!
    private var conn: ClientConnection!
    private var provider: TestBookServiceProvider!
    private var service: GRPCBookService!

    override func setUp() {
        serverGroup = PlatformSupport.makeEventLoopGroup(loopCount: 1)
        provider = TestBookServiceProvider()
        server = try! Server.start(configuration: Server.Configuration(target: .hostAndPort("localhost", 0), eventLoopGroup: serverGroup, serviceProviders: [provider])).wait()
        
        clientGroup = PlatformSupport.makeEventLoopGroup(loopCount: 1)
        conn = ClientConnection(configuration: ClientConnection.Configuration(target: .hostAndPort("localhost", server.channel.localAddress!.port!), eventLoopGroup: clientGroup))
        let client = Protobuf_BookServiceServiceClient(connection: conn)
        service = GRPCBookService(client: client)
    }

    override func tearDown() {
        try! conn.close().wait()
        try! clientGroup.syncShutdownGracefully()
        
        try! server.close().wait()
        try! serverGroup.syncShutdownGracefully()
    }
    
    func testList() {
        var tmp = Protobuf_Book()
        tmp.id = 1
        tmp.title = "abc"
        tmp.author = "xyz"
        let dto = [tmp]
        var response = Protobuf_BookListResponse()
        response.books = dto
        provider.listResponse = response
        let expectation = self.expectation(description: "Callback")
        service.list { result in
            switch result {
            case .success(let books):
                XCTAssertEqual(dto[0].id, Int64(books[0].id!))
                XCTAssertEqual(dto[0].title, books[0].title)
                XCTAssertEqual(dto[0].author, books[0].author)
            case .failure(let error):
                XCTFail(error.localizedDescription)
            }
            expectation.fulfill()
        }
        wait(for: [expectation], timeout: 1.0)
    }
    
    func testGet() {
        let id = 1
        var dto = Protobuf_Book()
        dto.id = Int64(id)
        dto.title = "abc"
        dto.author = "xyz"
        var response = Protobuf_BookGetResponse()
        response.book = dto
        provider.getResponse = response
        let expectation = self.expectation(description: "Callback")
        service.get(id: id) { result in
            switch result {
            case .success(let book):
                XCTAssertEqual(Int64(id), self.provider.getRequest.id)
                XCTAssertEqual(dto.id, Int64(book.id!))
                XCTAssertEqual(dto.title, book.title)
                XCTAssertEqual(dto.author, book.author)
            case .failure(let error):
                XCTFail(error.localizedDescription)
            }
            expectation.fulfill()
        }
        wait(for: [expectation], timeout: 1.0)
    }
    
    func testCreate() {
        let book = Book(id: nil, title: "abc", author: "xyz")
        provider.createResponse = Protobuf_BookCreateResponse()
        let expectation = self.expectation(description: "Callback")
        service.create(book: book) { result in
            switch result {
            case .success:
                XCTAssertEqual(book.title, self.provider.createRequest.book.title)
                XCTAssertEqual(book.author, self.provider.createRequest.book.author)
            case .failure(let error):
                XCTFail(error.localizedDescription)
            }
            expectation.fulfill()
        }
        wait(for: [expectation], timeout: 1.0)
    }
    
    func testUpdate() {
        let id = 1
        let book = Book(id: nil, title: "abc", author: "xyz")
        provider.updateResponse = Protobuf_BookUpdateResponse()
        let expectation = self.expectation(description: "Callback")
        service.update(id: id, book: book) { result in
            switch result {
            case .success:
                XCTAssertEqual(Int64(id), self.provider.updateRequest.id)
                XCTAssertEqual(book.title, self.provider.updateRequest.book.title)
                XCTAssertEqual(book.author, self.provider.updateRequest.book.author)
            case .failure(let error):
                XCTFail(error.localizedDescription)
            }
            expectation.fulfill()
        }
        wait(for: [expectation], timeout: 1.0)
    }
    
    func testDelete() {
        let id = 1
        provider.deleteResponse = Protobuf_BookDeleteResponse()
        let expectation = self.expectation(description: "Callback")
        service.delete(id: id) { result in
            switch result {
            case .success:
                XCTAssertEqual(Int64(id), self.provider.deleteRequest.id)
            case .failure(let error):
                XCTFail(error.localizedDescription)
            }
            expectation.fulfill()
        }
        wait(for: [expectation], timeout: 1.0)
    }
    
    class TestBookServiceProvider: Protobuf_BookServiceProvider {
        
        var listRequest: Protobuf_BookListRequest!
        var listResponse: Protobuf_BookListResponse!
        var createRequest: Protobuf_BookCreateRequest!
        var createResponse: Protobuf_BookCreateResponse!
        var getRequest: Protobuf_BookGetRequest!
        var getResponse: Protobuf_BookGetResponse!
        var updateRequest: Protobuf_BookUpdateRequest!
        var updateResponse: Protobuf_BookUpdateResponse!
        var deleteRequest: Protobuf_BookDeleteRequest!
        var deleteResponse: Protobuf_BookDeleteResponse!
        
        func list(request: Protobuf_BookListRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Protobuf_BookListResponse> {
            listRequest = request
            return context.eventLoop.makeSucceededFuture(listResponse)
        }
        
        func create(request: Protobuf_BookCreateRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Protobuf_BookCreateResponse> {
            createRequest = request
            return context.eventLoop.makeSucceededFuture(createResponse)
        }
        
        func get(request: Protobuf_BookGetRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Protobuf_BookGetResponse> {
            getRequest = request
            return context.eventLoop.makeSucceededFuture(getResponse)
        }
        
        func update(request: Protobuf_BookUpdateRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Protobuf_BookUpdateResponse> {
            updateRequest = request
            return context.eventLoop.makeSucceededFuture(updateResponse)
        }
        
        func delete(request: Protobuf_BookDeleteRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Protobuf_BookDeleteResponse> {
            deleteRequest = request
            return context.eventLoop.makeSucceededFuture(deleteResponse)
        }
        
    }

}

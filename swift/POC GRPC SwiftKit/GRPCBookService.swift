//
//  GRPCBookService.swift
//  POC GRPC SwiftKit
//
//  Created by Christopher San Diego on 1/15/20.
//  Copyright © 2020 Christopher San Diego. All rights reserved.
//

public class GRPCBookService {
    
    private let client: Protobuf_BookServiceServiceClient
    
    public init(client: Protobuf_BookServiceServiceClient) {
        self.client = client
    }
    
    public func list(completion: @escaping (Result<[Book], Error>) -> Void) {
        client.list(Protobuf_BookListRequest()).response.map { response in
            response.books.map { dto in
                Book(id: Int(dto.id), title: dto.title, author: dto.author)
            }
        }.whenComplete { result in
            completion(result)
        }
    }
    
    public func create(book: Book, completion: @escaping (Result<Bool?, Error>) -> Void) {
        var dto = Protobuf_Book()
        dto.title = book.title
        dto.author = book.author
        var request = Protobuf_BookCreateRequest()
        request.book = dto
        client.create(request).response.map { _ in
            nil as Bool?
        }.whenComplete { result in
            completion(result)
        }
    }
    
    public func get(id: Int, completion: @escaping (Result<Book, Error>) -> Void) {
        var request = Protobuf_BookGetRequest()
        request.id = Int64(id)
        client.get(request).response.map { response in
            Book(id: Int(response.book.id), title: response.book.title, author: response.book.author)
        }.whenComplete { result in
            completion(result)
        }
    }
    
    public func update(id: Int, book: Book, completion: @escaping (Result<Bool?, Error>) -> Void) {
        var dto = Protobuf_Book()
        dto.title = book.title
        dto.author = book.author
        var request = Protobuf_BookUpdateRequest()
        request.id = Int64(id)
        request.book = dto
        client.update(request).response.map { _ in
            nil as Bool?
        }.whenComplete { result in
            completion(result)
        }
    }
    
    public func delete(id: Int, completion: @escaping (Result<Bool?, Error>) -> Void) {
        var request = Protobuf_BookDeleteRequest()
        request.id = Int64(id)
        client.delete(request).response.map { _ in
            nil as Bool?
        }.whenComplete { result in
            completion(result)
        }
    }
}

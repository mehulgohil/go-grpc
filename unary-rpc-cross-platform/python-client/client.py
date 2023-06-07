import grpc
import transaction_pb2_grpc as pb2_grpc
import transaction_pb2 as pb2


def make_transaction():
    channel = grpc.insecure_channel(
        'localhost:8000')
    stub = pb2_grpc.MoneyTransactionStub(channel)

    request = pb2.TransactionRequest()
    request.PaidBy = 'Mehul'
    request.PaidTo = 'Hiren'
    request.Amount = 100.0

    response = stub.MakeTransaction(request)
    print('Transaction confirmation:', response.Confirmation)


if __name__ == '__main__':
    make_transaction()

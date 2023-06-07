from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class TransactionRequest(_message.Message):
    __slots__ = ["Amount", "PaidBy", "PaidTo"]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    Amount: float
    PAIDBY_FIELD_NUMBER: _ClassVar[int]
    PAIDTO_FIELD_NUMBER: _ClassVar[int]
    PaidBy: str
    PaidTo: str
    def __init__(self, PaidBy: _Optional[str] = ..., PaidTo: _Optional[str] = ..., Amount: _Optional[float] = ...) -> None: ...

class TransactionResponse(_message.Message):
    __slots__ = ["Confirmation"]
    CONFIRMATION_FIELD_NUMBER: _ClassVar[int]
    Confirmation: bool
    def __init__(self, Confirmation: bool = ...) -> None: ...

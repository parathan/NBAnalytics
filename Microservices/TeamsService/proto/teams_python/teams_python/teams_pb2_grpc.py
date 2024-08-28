# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from . import teams_pb2 as teams__pb2

GRPC_GENERATED_VERSION = '1.63.0'
GRPC_VERSION = grpc.__version__
EXPECTED_ERROR_RELEASE = '1.65.0'
SCHEDULED_RELEASE_DATE = 'June 25, 2024'
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    warnings.warn(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in teams_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class TeamsServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetTwoTeams = channel.unary_unary(
                '/teams.TeamsService/GetTwoTeams',
                request_serializer=teams__pb2.TwoTeamsRequest.SerializeToString,
                response_deserializer=teams__pb2.TwoTeamsResponse.FromString,
                _registered_method=True)
        self.GetAllTeams = channel.unary_unary(
                '/teams.TeamsService/GetAllTeams',
                request_serializer=teams__pb2.AllTeamsRequest.SerializeToString,
                response_deserializer=teams__pb2.AllTeamsResponse.FromString,
                _registered_method=True)
        self.GetTwoTeamsOrdered = channel.unary_unary(
                '/teams.TeamsService/GetTwoTeamsOrdered',
                request_serializer=teams__pb2.TwoTeamsRequest.SerializeToString,
                response_deserializer=teams__pb2.TwoTeamsOrderedResponse.FromString,
                _registered_method=True)


class TeamsServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetTwoTeams(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAllTeams(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetTwoTeamsOrdered(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TeamsServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetTwoTeams': grpc.unary_unary_rpc_method_handler(
                    servicer.GetTwoTeams,
                    request_deserializer=teams__pb2.TwoTeamsRequest.FromString,
                    response_serializer=teams__pb2.TwoTeamsResponse.SerializeToString,
            ),
            'GetAllTeams': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAllTeams,
                    request_deserializer=teams__pb2.AllTeamsRequest.FromString,
                    response_serializer=teams__pb2.AllTeamsResponse.SerializeToString,
            ),
            'GetTwoTeamsOrdered': grpc.unary_unary_rpc_method_handler(
                    servicer.GetTwoTeamsOrdered,
                    request_deserializer=teams__pb2.TwoTeamsRequest.FromString,
                    response_serializer=teams__pb2.TwoTeamsOrderedResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'teams.TeamsService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class TeamsService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetTwoTeams(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/teams.TeamsService/GetTwoTeams',
            teams__pb2.TwoTeamsRequest.SerializeToString,
            teams__pb2.TwoTeamsResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetAllTeams(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/teams.TeamsService/GetAllTeams',
            teams__pb2.AllTeamsRequest.SerializeToString,
            teams__pb2.AllTeamsResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetTwoTeamsOrdered(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/teams.TeamsService/GetTwoTeamsOrdered',
            teams__pb2.TwoTeamsRequest.SerializeToString,
            teams__pb2.TwoTeamsOrderedResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

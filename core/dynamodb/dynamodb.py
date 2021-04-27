
import boto3


class DynamoClient:
    def __init__(self):
        self.client = boto3.client('dynamodb', region_name="us-east-1")

    def GetAllTables(self):
        response = self.client.list_tables(
            Limit=30
        )
        return response["TableNames"]

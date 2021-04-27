from django.shortcuts import render

from django.http import HttpResponse
from django.http import JsonResponse

from core.dynamodb.dynamodb import DynamoClient

def dynamodb_tables_ui(request):
    dynamoClient = DynamoClient()
    result = dynamoClient.GetAllTables()
    return HttpResponse("Dynamodb tables")

def dynamodb_tables_data(request):
    dynamoClient = DynamoClient()
    result = dynamoClient.GetAllTables()
    return JsonResponse({
        "result": result
    })
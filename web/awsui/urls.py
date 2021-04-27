from django.urls import path

from . import views
from . import dynamodb

urlpatterns = [
    path('', views.index, name='index'),
    path('dynamodb/ui/tables', dynamodb.dynamodb_tables_ui, name='index'),
    path('dynamodb/data/tables', dynamodb.dynamodb_tables_data, name='index'),
]
from rest_framework import viewsets
from .serializer import TaskSerializer
from .models import Task

# Create your views here.
class TaskView(viewsets.ModelViewSet):
    serializer_class = TaskSerializer # This is the serializer class
    queryset = Task.objects.all() # This is the queryset


from django.db import models

# Create your models here.
class Task (models.Model):
    title = models.CharField(max_length=200)
    description = models.TextField(blank=True)
    completed = models.BooleanField(default=False)

    # This is the string representation of the model
    def __str__(self):
        return self.title

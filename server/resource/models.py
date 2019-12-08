from django.db import models
from django.contrib.auth.models import User

# Create your models here.
class Resource(models.Model):
    PLATFORM = (
        ('1', 'vsphere'),
        ('2', 'kvm'),
        ('3', 'hyper-v')
    )

    STATUS = (
        ('1', 'idle'),
        ('2', 'normal'),
        ('3', 'busy')
    )

    # name = models.CharField(max_length=200)
    date_added = models.DateTimeField(blank=True, null=True)
    host_address = models.TextField()
    host_name = models.TextField()
    platform_type = models.CharField(max_length=10, choices=PLATFORM, default='vsphere')
    username = models.CharField(unique=True, max_length=150)
    password = models.CharField(max_length=128)
    status = models.CharField(max_length=20, choices=STATUS, default='normal')
    total_disk = models.FloatField(blank=True, null=True)
    total_ram = models.FloatField(blank=True, null=True)
    total_cpu = models.FloatField(blank=True, null=True)
    current_disk = models.FloatField(blank=True, null=True)
    current_ram = models.FloatField(blank=True, null=True)
    current_cpu = models.FloatField(blank=True, null=True)
    is_active = models.BooleanField(default=False)
    total_jobs = models.IntegerField(blank=True, null=True, default=0)
    job_completed = models.IntegerField(blank=True, null=True, default=0)
    polling_interval = models.IntegerField(blank=True, null=True)
    monitored = models.BooleanField(blank=True, null=True, default=False)
    user = models.ManyToManyField(User, blank=True)

    class Meta:
        verbose_name = 'Tides Resource'
        verbose_name_plural = 'Tides Resources'

    def save(self, *args, **kwargs):
        # do something
        super().save(*args, **kwargs)
        # do something

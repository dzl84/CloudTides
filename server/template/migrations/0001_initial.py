# Generated by Django 2.2.7 on 2019-12-08 03:21

from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Template',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=150, unique=True)),
                ('date_added', models.DateTimeField(blank=True, null=True)),
                ('guest_os', models.CharField(max_length=100)),
                ('compatibility', models.CharField(blank=True, max_length=100, null=True)),
                ('provisioned_space', models.FloatField(blank=True, null=True)),
                ('memory_size', models.FloatField(blank=True, null=True)),
                ('template_type', models.CharField(choices=[('1', 'host'), ('2', 'tides')], default='tides', max_length=20)),
                ('username', models.TextField()),
                ('password', models.TextField()),
            ],
            options={
                'verbose_name': 'Tides Template',
                'verbose_name_plural': 'Tides Templates',
            },
        ),
    ]
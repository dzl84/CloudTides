from django.shortcuts import render
from rest_framework.response import Response
from rest_framework.views import APIView
import json
import datetime
from .models import Policy
from django.shortcuts import get_object_or_404
from resource.models import *
from django.contrib.auth.models import User
from rest_framework.permissions import IsAuthenticated
from rest_framework.authtoken.models import Token
from .serializers import *
from projects.models import *
from django.utils import timezone
from django.contrib.auth.hashers import *
from template.models import *

# Create your views here.

class AddPolicy(APIView):
    permission_classes = (IsAuthenticated,)

    def post(self, request):
        data = json.loads(request.body)
        # password = make_password(data['password'])
        password = data['password']
        username = data['username']
        account_type = data['accountType']
        name = data['name']
        is_destroy = None
        if data['is_destroy'].lower() == "true":
            is_destroy = True
        elif data['is_destroy'].lower() == "false":
            is_destroy = False
        deploy_type = data['deployType']
        idle_policy = data['idle']
        threshold_policy = data['threshold']
        project_id = data['project']
        template_id = data['template']
        date_created = timezone.make_aware(datetime.datetime.now(), timezone.get_default_timezone())

        token = request.META.get('HTTP_AUTHORIZATION').split(' ')[1]
        user = Token.objects.get(key=token).user

        project = Projects.objects.get(pk=project_id)
        template = Template.objects.get(pk=template_id)
        policy = Policy(username=username, password=password, name=name, date_created=date_created,
                        is_destroy=is_destroy, deploy_type=deploy_type, idle_policy=idle_policy,
                        account_type=account_type, threshold_policy=threshold_policy,
                        project=project, template=template, user=user)
        #try:
        policy.save()
        #policy.user.set((user,))
        #except:
            #return Response({'message: policy name exists'}, status=401)

        return Response({'message': 'success'}, status=200)


class UpdatePolicy(APIView):
    permission_classes = (IsAuthenticated,)

    def post(self, request):
        data = json.loads(request.body)
        password = data['password']
        username = data['username']
        account_type = data['accountType']
        name = data['name']
        is_destroy = None
        if data['is_destroy'].lower() == "true":
            is_destroy = True
        elif data['is_destroy'].lower() == "false":
            is_destroy = False
        deploy_type = data['deployType']
        idle_policy = data['idle']
        threshold_policy = data['threshold']
        project_id = data['project']
        template_id = data['template']
        date_created = timezone.make_aware(datetime.datetime.now(), timezone.get_default_timezone())

        token = request.META.get('HTTP_AUTHORIZATION').split(' ')[1]
        user = Token.objects.get(key=token).user

        project = Projects.objects.get(pk=project_id)
        template = Template.objects.get(pk=template_id)

        obj = get_object_or_404(Policy, name=name)
        obj.username = username
        obj.password = password
        obj.account_type = account_type
        obj.is_destroy = is_destroy
        obj.deploy_type = deploy_type
        obj.idle_policy = idle_policy
        obj.threshold_policy = threshold_policy
        obj.user = user
        obj.project = project
        obj.template = template
        obj.date_created = date_created
    
        obj.save(update_fields=['is_destroy', 'deploy_type', 'idle_policy', 'user', 'username', 'password', 'account_type',
                        'threshold_policy', 'project', 'template', 'date_created'])

        return Response({'message': 'success'}, status=200)


class ListPolicy(APIView):
    permission_classes = (IsAuthenticated,)

    def get(self, request):
        token = request.META.get('HTTP_AUTHORIZATION').split(' ')[1]
        user = Token.objects.get(key=token).user
        policies = Policy.objects.filter(user=user)
        results = []
        for policy in policies:
            hosts_number = Resource.objects.filter(policy=policy).count()
            data = PolicySerializer(policy).data
            data["hosts_assigned"] = hosts_number
            results.append(data)

        return Response({'message': 'success', 'status': True, 'results': results}, status=200)


class RemovePolicy(APIView):

    def post(self, request):
        data = json.loads(request.body)
        host_address = data['host_address']
        name = data['name']

        try:
            Policy.objects.get(host_address=host_address, name=name).delete()
        except:
            return Response({'message': 'no matching objects'}, status=401)

        return Response({'message': 'success'}, status=200)

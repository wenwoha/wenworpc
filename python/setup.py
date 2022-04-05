#!/usr/bin/env python
# -*- encoding: utf-8 -*-

from setuptools import setup, find_packages

setup(
    name='wenworpc',
    version='0.1.3',
    description='internal gRPC protos and stubs',
    author='wenwoha',
    author_email='guoshijiang2012@163.com',
    packages=find_packages(),
    install_requires=[
        'grpcio>=1.27.2',
        'grpcio-tools>=1.27.2',
        'protobuf>=3.11.3',
        'PyYAML>=5.3.1'
    ],
    python_requires='>=2.7.12',
)
from setuptools import setup
import os, sys

def read(fname):
    return open(os.path.join(os.path.dirname(__file__), fname)).read()

if sys.version_info < (3,4):
	sys.exit('Python >= 3.4 is required')

setup(name='StrongPasswordGenerator',
  version='0.1',
  description='Generates a strong password',
  long_description=read('README.rst'),
  url='http://github.com/ogelami/StrongPasswordGenerator',
  author='ogelami',
  author_email='ogelami@gmail.com',
  scripts=['bin/strongPasswordGenerator'],
  license='MIT',
  packages=['StrongPasswordGenerator'],
  zip_safe=False)

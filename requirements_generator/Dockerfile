# Dockerfile for requirements_generator
FROM python:3.9-slim

WORKDIR /usr/src/app

COPY requirements_generator/requirements.txt ./requirements/
COPY requirements_generator/main.py ./requirements/
COPY requirements_generator/create/go_struct_generator.py ./requirements/create/
COPY requirements_generator/create/xsd_processor.py ./requirements/create/

RUN pip install --no-cache-dir -r requirements/requirements.txt

CMD ["python", "./requirements/main.py"]

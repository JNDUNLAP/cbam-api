
FROM python:3.9-slim

WORKDIR /usr/src/app

COPY requirements.txt ./requirements/
COPY main.py ./requirements/
COPY create/go_struct_generator.py ./requirements/create/
COPY create/xsd_processor.py ./requirements/create/


RUN pip install --no-cache-dir -r requirements/requirements.txt

CMD ["python", "./requirements/main.py"]

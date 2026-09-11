FROM python:3.11-slim

# Latest upstream commit (2023-06-06); the v3.0.0 tag dates from 2020.
ARG MEDALLION_REF=39e76bf18be5371e9570de7e5f340c3937b69c0d

# setuptools 82 removed pkg_resources, which medallion still imports.
RUN pip install --no-cache-dir setuptools==81.0.0 \
    https://github.com/oasis-open/cti-taxii-server/archive/${MEDALLION_REF}.tar.gz

COPY config.json default_data.json /opt/taxii2/

EXPOSE 5000

# The Flask dev server ignores SIGTERM as PID 1 but exits cleanly on SIGINT.
STOPSIGNAL SIGINT

HEALTHCHECK --interval=5s --timeout=3s --start-period=5s --retries=10 \
    CMD ["python", "-c", "import urllib.request as r; r.urlopen(r.Request('http://127.0.0.1:5000/taxii2/', headers={'Accept': 'application/taxii+json;version=2.1', 'Authorization': 'Basic YWRtaW46UGFzc3dvcmQwMQ=='}), timeout=2)"]

CMD ["medallion", "--host", "0.0.0.0", "--port", "5000", "--log-level=DEBUG", "/opt/taxii2/config.json"]

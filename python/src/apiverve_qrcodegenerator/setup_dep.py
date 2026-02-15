from setuptools import setup, find_packages

setup(
    name='apiverve_qrcodegenerator',
    version='1.1.13',
    packages=find_packages(),
    include_package_data=True,
    install_requires=[
        'requests',
        'setuptools'
    ],
    description='QR Code Generator creates customizable QR codes with support for colors, gradients, logos, and various styling options. Generate professional QR codes for marketing, packaging, and digital experiences.',
    author='APIVerve',
    author_email='hello@apiverve.com',
    url='https://apiverve.com/marketplace/qrcodegenerator?utm_source=pypi&utm_medium=homepage',
    classifiers=[
        'Programming Language :: Python :: 3',
        'Operating System :: OS Independent',
    ],
    python_requires='>=3.6',
)

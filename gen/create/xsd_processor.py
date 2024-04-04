from lxml import etree
import os

namespace = {'xs': 'http://www.w3.org/2001/XMLSchema'}

def extract_simple_types(xsd_tree):
    simple_types = {}
    for simple_type in xsd_tree.xpath('//xs:simpleType', namespaces=namespace):
        type_name = simple_type.get('name')
        restrictions = simple_type.xpath('.//xs:restriction', namespaces=namespace)[0]
        base_type = restrictions.get('base')
        
        restriction_details = {}
        for restriction in restrictions.getchildren():
            restriction_tag = etree.QName(restriction).localname
            restriction_value = restriction.get('value')
            restriction_details[restriction_tag] = restriction_value
        
        simple_types[type_name] = {
            'base_type': base_type,
            'restrictions': restriction_details
        }
    return simple_types

def link_elements_to_types(xsd_tree, simple_types):
    element_details = {}
    elements = xsd_tree.xpath('//xs:element', namespaces=namespace)
    for element in elements:
        element_name = element.get('name')
        element_type = element.get('type')
        
        if element_type in simple_types:
            element_details[element_name] = {
                'type': element_type,
                'restrictions': simple_types[element_type]['restrictions']
            }
    return element_details

def process_xsd_files(file_list):
    all_simple_types = {}
    all_element_details = {}
    for xsd_file in file_list:
        if not os.path.isfile(xsd_file):
            print(f"File not found: {xsd_file}")
            continue
        
        tree = etree.parse(xsd_file)
        simple_types_info = extract_simple_types(tree)
        all_simple_types.update(simple_types_info)
    
    for xsd_file in file_list:
        if 'stypes' in xsd_file:  # Skip the simple types file for element linking
            continue
        tree = etree.parse(xsd_file)
        element_details = link_elements_to_types(tree, all_simple_types)
        all_element_details.update(element_details)
    
    return all_element_details

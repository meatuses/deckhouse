# coding: utf-8

"""
    Kubernetes

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: release-1.25
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from kubernetes.client.configuration import Configuration


class V1ReplicaSetSpec(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'min_ready_seconds': 'int',
        'replicas': 'int',
        'selector': 'V1LabelSelector',
        'template': 'V1PodTemplateSpec'
    }

    attribute_map = {
        'min_ready_seconds': 'minReadySeconds',
        'replicas': 'replicas',
        'selector': 'selector',
        'template': 'template'
    }

    def __init__(self, min_ready_seconds=None, replicas=None, selector=None, template=None, local_vars_configuration=None):  # noqa: E501
        """V1ReplicaSetSpec - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._min_ready_seconds = None
        self._replicas = None
        self._selector = None
        self._template = None
        self.discriminator = None

        if min_ready_seconds is not None:
            self.min_ready_seconds = min_ready_seconds
        if replicas is not None:
            self.replicas = replicas
        self.selector = selector
        if template is not None:
            self.template = template

    @property
    def min_ready_seconds(self):
        """Gets the min_ready_seconds of this V1ReplicaSetSpec.  # noqa: E501

        Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)  # noqa: E501

        :return: The min_ready_seconds of this V1ReplicaSetSpec.  # noqa: E501
        :rtype: int
        """
        return self._min_ready_seconds

    @min_ready_seconds.setter
    def min_ready_seconds(self, min_ready_seconds):
        """Sets the min_ready_seconds of this V1ReplicaSetSpec.

        Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)  # noqa: E501

        :param min_ready_seconds: The min_ready_seconds of this V1ReplicaSetSpec.  # noqa: E501
        :type: int
        """

        self._min_ready_seconds = min_ready_seconds

    @property
    def replicas(self):
        """Gets the replicas of this V1ReplicaSetSpec.  # noqa: E501

        Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller  # noqa: E501

        :return: The replicas of this V1ReplicaSetSpec.  # noqa: E501
        :rtype: int
        """
        return self._replicas

    @replicas.setter
    def replicas(self, replicas):
        """Sets the replicas of this V1ReplicaSetSpec.

        Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller  # noqa: E501

        :param replicas: The replicas of this V1ReplicaSetSpec.  # noqa: E501
        :type: int
        """

        self._replicas = replicas

    @property
    def selector(self):
        """Gets the selector of this V1ReplicaSetSpec.  # noqa: E501


        :return: The selector of this V1ReplicaSetSpec.  # noqa: E501
        :rtype: V1LabelSelector
        """
        return self._selector

    @selector.setter
    def selector(self, selector):
        """Sets the selector of this V1ReplicaSetSpec.


        :param selector: The selector of this V1ReplicaSetSpec.  # noqa: E501
        :type: V1LabelSelector
        """
        if self.local_vars_configuration.client_side_validation and selector is None:  # noqa: E501
            raise ValueError("Invalid value for `selector`, must not be `None`")  # noqa: E501

        self._selector = selector

    @property
    def template(self):
        """Gets the template of this V1ReplicaSetSpec.  # noqa: E501


        :return: The template of this V1ReplicaSetSpec.  # noqa: E501
        :rtype: V1PodTemplateSpec
        """
        return self._template

    @template.setter
    def template(self, template):
        """Sets the template of this V1ReplicaSetSpec.


        :param template: The template of this V1ReplicaSetSpec.  # noqa: E501
        :type: V1PodTemplateSpec
        """

        self._template = template

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1ReplicaSetSpec):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1ReplicaSetSpec):
            return True

        return self.to_dict() != other.to_dict()

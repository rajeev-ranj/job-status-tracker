import streamlit as st
import requests

def display_services():
    st.subheader("Services")

    # Fetch all services
    response = requests.get("http://localhost:8080/services")
    if response.status_code == 200:
        services = response.json()
    else:
        st.error("Failed to fetch services.")
        return

    # Display all services in a table
    st.write("All Services:")
    st.write(services)

    # Allow the user to select a service to view its details
    service_id = st.selectbox("Select a service to view its details:", [service['service_id'] for service in services])
    display_service_details(service_id)

def display_service_details(service_id):
    # Fetch the details for the selected service
    response = requests.get(f"http://localhost:8080/services/{service_id}")
    if response.status_code == 200:
        service_details = response.json()
    else:
        st.error(f"Failed to fetch details for service {service_id}.")
        return

    # Display the service details
    st.write(f"Details for Service {service_id}:")
    st.write(service_details)

    # Additional code to manage the service, e.g., update, delete, etc.
    # ...

    # Display jobs related to this service
    display_service_jobs(service_id)

def display_service_jobs(service_id):
    # Fetch the jobs for the selected service
    response = requests.get(f"http://localhost:8080/services/{service_id}/jobs")
    if response.status_code == 200:
        jobs = response.json()
    else:
        st.error(f"Failed to fetch jobs for service {service_id}.")
        return

    # Display the jobs for the selected service
    st.write(f"Jobs for Service {service_id}:")
    st.write(jobs)

    # Additional code to manage the jobs, e.g., create, update, delete, etc.
    # ...

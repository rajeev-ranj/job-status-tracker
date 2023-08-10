import streamlit as st
import requests

# URL of the backend API
API_URL = "http://localhost:8080"

def show_dashboard():
    st.title("Central Dashboard")
    st.write("Here you can view all the jobs and their history for the last few days.")

    # Fetching the jobs from the backend
    response = requests.get(f"{API_URL}/jobs")
    if response.status_code == 200:
        jobs = response.json()
        display_jobs(jobs)
    else:
        st.error("Failed to fetch jobs from the backend.")

def display_jobs(jobs):
    # Displaying the jobs in a table
    st.write("### Jobs Overview")
    st.table(jobs)

    # Displaying the history for each job
    for job in jobs:
        st.write(f"### Job: {job['name']}")
        display_job_history(job['id'])

def display_job_history(job_id):
    # Fetching the job history from the backend
    response = requests.get(f"{API_URL}/jobs/{job_id}/history")
    if response.status_code == 200:
        history = response.json()
        st.table(history)
    else:
        st.error(f"Failed to fetch history for job ID {job_id}.")

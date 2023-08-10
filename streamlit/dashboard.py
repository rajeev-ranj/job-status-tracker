import streamlit as st
import requests

def display_dashboard():
    st.subheader("Central Dashboard")

    # Fetch all jobs
    response = requests.get("http://localhost:8080/jobs")
    if response.status_code == 200:
        jobs = response.json()
    else:
        st.error("Failed to fetch jobs.")
        return

    # Display all jobs in a table
    st.write("All Jobs:")
    st.write(jobs)

    # Allow the user to select a job to view its history
    job_id = st.selectbox("Select a job to view its history:", [job['job_id'] for job in jobs])
    display_job_history(job_id)

def display_job_history(job_id):
    # Fetch the history for the selected job
    response = requests.get(f"http://localhost:8080/jobs/{job_id}/history")
    if response.status_code == 200:
        history = response.json()
    else:
        st.error(f"Failed to fetch history for job {job_id}.")
        return

    # Display the job history
    st.write(f"History for Job {job_id}:")
    st.write(history)

    # Additional code to visualize the history, e.g., using charts or graphs
    # ...

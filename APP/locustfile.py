from locust import HttpLocust, TaskSet, task

class UserBehavior(TaskSet):

    @task(2)
    def index(self):
        self.client.get("/")

    @task(1)
    def profile(self):
        self.client.get("/getData")

class WebsiteUser(HttpLocust):
    task_set = UserBehavior
    
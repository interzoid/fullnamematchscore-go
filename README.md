## fullnamematchscore-go

Go package for generating a matching score of two individual person names from 0-100, where 100 is the highest matching score, on how closely two individual full names match. The scoring is based on a series of tests, algorithms, AI, and an ever-growing body of Machine Learning-based generated knowledge.

### Usage

To generate the match score, you will need the following information:

- an API License key, available at https://www.interzoid.com
- two individual full names to compare

Begin by importing the package:

    import "github.com/interzoid/fullnamematchscore-go"

Then, feed the information into the GetScore() method:

    score, code, credits, err := FullNameMatchScore.GetScore("YOUR-API-KEY","Jim Smith","Mr. James Smyth")


The return values will be the generated match comparison score (0-100), a code (success or failure), how many remaining credits on your API key, and any error messages. The score allows you to set a score threshold in your own logic for your specific case, for example, any score higher than 50 can be considered a "match" (or 60, 70, etc.)
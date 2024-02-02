import React from "react";
import "./ImagesOptions.css"

const ImagesOptions = ({ subtext }) => {
  const regex = /!\[branch-1\]\(([^)]+)\)/g;
  const imagesMatches = [...subtext.matchAll(regex)];
  let text = subtext;

  imagesMatches.forEach((element) => {
    text = text.replace(element[0], "");
  });

  const imagesUrl = imagesMatches.map((match) => match[1]);
  text = text.split("\n\n");
  return (
    <div>
      <p>{text[0]}</p>
      {imagesUrl.map((image, index) => {
        return (
          <figure className="figure-option">
            <img className="figure-option__image" src={image} alt={`option ${text[index + 1]}`} />
            <figcaption className="figure-option__figcaption">{text[index + 1]}</figcaption>
          </figure>
        );
      })}
    </div>
  );
};

export default ImagesOptions;
